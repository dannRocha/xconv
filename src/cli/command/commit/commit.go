package commit

import (
	"fmt"
	"strings"

	"github.com/dannrocha/czen/src/cli"
	"github.com/dannrocha/czen/src/gitscm"
	"github.com/dannrocha/czen/src/setup"
	"github.com/manifoldco/promptui"
)

func Execute(args ...string) error {
	messages := make(map[string]string)

	role := setup.Role{}
	scrip := setup.Script{}

	role.LoadRole()
	scrip.LoadScript()

	profile, profileErr := role.FindCurrentProfileEnable()

	for _, auto := range scrip.Automation {

		if auto.Bind == cmd.COMMIT && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	if profileErr != nil {
		return nil
	}

	var questionGroup setup.Question = setup.Question{}

	for _, question := range profile.Questions {
		if question.Type == setup.LIST {
			questionGroup = question
			break
		}
	}

	options := questionGroup.Choices
	sampleInputGroup := []setup.Question{}
	listInputGroup := setup.Question{}

	for _, question := range profile.Questions {
		if question.Type == setup.LIST {
			listInputGroup = question
			continue
		}

		sampleInputGroup = append(sampleInputGroup, question)
	}

	prompt := promptui.Select{
		Label: listInputGroup.Message,
		Items: optionDescription(options),
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	messages[questionGroup.Name] = options[index].Value

	parse := map[string]func(string) string{

		setup.INPUT: func(content string) string {
			return strings.TrimSpace(content)
		},

		setup.CONFIRM: func(content string) string {
			trueValues := []string{
				"y",
				"yes",
			}

			var value bool = false
			content = strings.ToLower(strings.TrimSpace(content))

			for _, confirm := range trueValues {
				if confirm == content {
					value = true
					break
				}
			}

			return fmt.Sprintf("%t", value)
		},
	}

	for _, input := range sampleInputGroup {
		clientInput := ""

		if input.Type == setup.CONFIRM {
			confirm := promptui.Prompt{
				Label:     input.Message,
				Default:   input.Default,
				IsConfirm: true,
			}

			action, err := confirm.Run()

			if err != nil {
				action = fmt.Sprint(false)
			}

			clientInput = action

		} else {
			for {

				clientInput = parse[input.Type](cmd.ReadInput(input.Message))

				filter, isContaisFilter := scrip.FindByFilterName(input.Filter)

				for _, middlewareName := range input.Middleware {

					middleware, isContaisMiddleware := scrip.FindByMiddlewareName(middlewareName)

					if isContaisMiddleware {
						if middleware.Enable {
							clientInput = middleware.Run(clientInput)
						}
					}
				}

				if isContaisFilter {

					if filter.Enable {

						if !filter.Run(clientInput) {
							break
						} else {

							message := filter.Message

							if message.Content != "" {
								if message.Color {
									fmt.Printf("\033[33m%v\033[0m\n", filter.Message.Content)
								} else {
									fmt.Printf("%v\n", filter.Message.Content)
								}
							}
						}
					}

					if !filter.Retry || !filter.Enable {
						break
					}
				} else {
					break
				}
			}
		}

		if input.Default != "" && strings.TrimSpace(clientInput) == "" {
			messages[input.Name] = parse[input.Type](input.Default)
			continue
		}

		messages[input.Name] = clientInput
	}

	commitMessage, _ := gitscm.FormatCommit(messages)

	command := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"commit",
			"-m",
			commitMessage,
		},
	}

	stdout, err := command.Execute()

	if err != nil {
		fmt.Println(string(stdout))
	}

	return nil
}

func optionDescription(optionGroup []setup.Option) []string {

	names := []string{}

	for _, option := range optionGroup {
		names = append(names, option.Name)
	}

	return names
}
