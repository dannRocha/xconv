package application

import (
	"log"
	"os"
	"time"

	"github.com/dannrocha/xconv/src/cmd/cli"
	CLI "github.com/urfave/cli/v2"
)

func Run() {
	CLI.VersionFlag = &CLI.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &CLI.App{
		Name:     "xconv",
		Compiled: time.Now(),
		Authors: []*CLI.Author{
			{
				Name:  "Rocha da Silva, Daniel",
				Email: "rochadaniel@acad.ifma.edu.br",
			},
		},
		Copyright: "(c) 2022 MIT",
		HelpName:  "contrive",
		Usage:     "XConventional is a cli tool to generate conventional commits.",
		UsageText: "xconv [-h] {init,commit,example,info,tag,schema,bump,changelog,version}",
		ArgsUsage: "[args and such]",
		HideHelp:  false,
		Commands: []*CLI.Command{
			{
				Name:    cli.INIT,
				Aliases: []string{"i"},
				Action:  cli.Init,
				Usage:   "init xconv configuration",
			},
			{
				Name:    cli.COMMIT,
				Aliases: []string{"c"},
				Action:  cli.Commit,
				Usage:   "create new commit",
			},
			{
				Name:    cli.CHANGELOG,
				Aliases: []string{"ch"},
				Action:  cli.Changelog,
				Usage:   "generate changelog (note that it will overwrite existing file)",
			},
			{
				Name:    cli.BUMP,
				Aliases: []string{"b"},
				Action:  cli.Bump,
				Usage:   "bump semantic version based on the git log",
			},
			{
				Name:    cli.ROLLBACK,
				Aliases: []string{"r"},
				Action:  cli.Rollback,
				Usage:   "revert commit to a specific tag",
			},
			{
				Name:    cli.TAG,
				Aliases: []string{"t"},
				Action:  cli.Tag,
				Usage:   "show tags",
			},
			{
				Name:    cli.SCHEMA,
				Aliases: []string{"s"},
				Action:  cli.Schema,
				Usage:   "show commit schema",
			},
			{
				Name:    cli.EXAMPLE,
				Aliases: []string{"e"},
				Action:  cli.Example,
				Usage:   "show commit example",
			},
			{
				Name:    cli.VERSION,
				Aliases: []string{"v"},
				Action:  cli.Version,
				Usage:   "get the version of the installed xconv or the current project",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
