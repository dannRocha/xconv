package setup

const (
	LUA     = "lua"
	SH      = "sh"
	BEFORE  = "before"
	AFTER   = "after"
	LIST    = "list"
	CONFIRM = "confirm"
	INPUT   = "input"
)

type Option struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
	Name  string `yaml:"name"`
}

type Question struct {
	Type       string   `yaml:"type"`
	Message    string   `yaml:"message"`
	Name       string   `yaml:"name"`
	Default    string   `yaml:"default"`
	Middleware []string `yaml:"middleware"`
	Filter     string   `yaml:"filter"`
	Editor     bool     `yaml:"editor"`
	Choices    []Option `yaml:"choices"`
}

type ExitCode struct {
	Exception   string
	ExitCode    int
	Description string
}

type Configuration struct {
	Rule   Rule   `yaml:"rule"`
	Script Script `yaml:"script"`
}

type Rule struct {
	Version       string    `yaml:"version"`
	ActiveProfile string    `yaml:"active_profile"`
	ChangelogFile string    `yaml:"changelog_file"`
	Profiles      []Profile `yaml:"profiles"`
}

type Tag struct {
	Format string `yaml:"format"`
	Mode   string `yaml:"mode"`
}

type Profile struct {
	Name            string            `yaml:"name"`
	BumpMap         map[string]string `yaml:"bump_map"`
	BumpPattern     string            `yaml:"bump_pattern"`
	CommitParser    string            `yaml:"commit_parser"`
	ChangeTypeOrder []string          `yaml:"change_type_order"`
	Example         string            `yaml:"example"`
	MessageTemplate string            `yaml:"message_template"`
	Questions       []Question        `yaml:"questions"`
	Schema          string            `yaml:"schema"`
	Tag             Tag               `yaml:"tag"`
}

type Message struct {
	Content string `yaml:"content"`
	Color   bool   `yaml:"color"`
}

type ScriptBase struct {
	Name   string `yaml:"name"`
	Enable bool   `yaml:"enable"`
	Script string `yaml:"script"`
	Type   string `yaml:"type"`
}

type Filter struct {
	ScriptBase `yaml:",inline"`
	Retry      bool    `yaml:"retry"`
	Message    Message `yaml:"message"`
}

type Task struct {
	ScriptBase `yaml:",inline"`
	Bind       string `yaml:"bind"`
	Language   string `yaml:"language"`
	When       string `yaml:"when"`
}

type Middleware struct {
	ScriptBase `yaml:",inline"`
}

type Script struct {
	Filter     []Filter     `yaml:"filters"`
	Middleware []Middleware `yaml:"middlewares"`
	Task       []Task       `yaml:"tasks"`
}

// https://github.com/commitizen-tools/commitizen/blob/master/docs/exit_codes.md
var ExitCodeStardard map[string]ExitCode = map[string]ExitCode{
	"ExpectedExit": {
		Exception:   "ExpectedExit",
		ExitCode:    0,
		Description: "Expected exit",
	},

	"FileSetupExist": {
		Exception:   "FileSetupExist",
		ExitCode:    1,
		Description: "there is .xconv.yaml configuration in the repository",
	},

	"NotAGitProjectError": {
		Exception:   "NotAGitProjectError",
		ExitCode:    2,
		Description: "Not in a git project",
	},

	"NoCommitsFoundError": {
		Exception:   "NoCommitsFoundError",
		ExitCode:    3,
		Description: "No commit found",
	},

	"NoVersionSpecifiedError": {
		Exception:   "NoCommitsFoundError",
		ExitCode:    4,
		Description: "Version can not be found in configuration file [.xconv.yaml]",
	},

	"NoPermissionOnDir": {
		Exception:   "NoPermissionOnDir",
		ExitCode:    5,
		Description: ".xconv.yaml file cannot be create in the current directory",
	},

	"NothingToCommitError": {
		Exception:   "NothingToCommitError",
		ExitCode:    11,
		Description: "Nothing in staging to be committed",
	},

	"MissingConfigError": {
		Exception:   "MissingConfigError",
		ExitCode:    15,
		Description: "Configuration missed for .xconv.yaml",
	},

	"CurrentVersionNotFoundError": {
		Exception:   "CurrentVersionNotFoundError",
		ExitCode:    17,
		Description: "current version cannot be found in version_files",
	},

	"InvalidConfigurationError": {
		Exception:   "InvalidConfigurationError",
		ExitCode:    19,
		Description: "An error was found in the xconv configuration",
	},

	"NoneIncrementExit": {
		Exception: "InvalidConfigurationError",
		ExitCode:  21,
		Description: "	The commits found are not elegible to be bumped",
	},

	"InvalidScriptFilter": {
		Exception:   "InvalidScriptFilter",
		ExitCode:    22,
		Description: "An error was found in the script",
	},
}
