package commands

import (
	"fmt"

	"cillers-cli/config"
)

// CommandHelp handles the 'help' command
func CommandHelp(args []string, options map[string]bool) error {
	cfg := config.Get()
	
	helpText := fmt.Sprintf(`Cillers CLI version %s

Usage: cillers [command] [options]

Commands:
  new <name>        Create a new system with the specified name
  help              Show this help message
  version           Show the version number

Options:
  --verbose         Enable verbose output for debugging purposes

Examples:
  cillers new my-project
  cillers new my-project --verbose
  cillers help
  cillers version

For more information, please visit: %s
`, cfg.Version, cfg.DocumentationURL)

	fmt.Println(helpText)
	
	return nil
}
