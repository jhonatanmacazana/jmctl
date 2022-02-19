package commands

type JmctlCommand struct {
	Help HelpCommand `command:"help" description:"Print this help message"`

	Version func() `short:"v" long:"version" description:"Print the version of jmctl and exit"`
	Verbose bool   `long:"verbose" description:"Print API requests and responses"`

	Template TemplatesCommand `command:"template"      alias:"ts" description:"List available templates"`
	Script   TemplatesCommand `command:"script"      alias:"sc" description:"List available scripts"`
	Snippet  TemplatesCommand `command:"snippet"      alias:"sn" description:"List available snippets"`

	// Curl CurlCommand `command:"curl" alias:"c" description:"curl the api"`
	// Completion CompletionCommand `command:"completion" description:"generate shell completion code"`
}

var JmCtl JmctlCommand
