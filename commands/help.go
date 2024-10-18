package commands

import (
	"fmt"

	"cillers-cli/config"
)

func Help(args []string, options map[string]bool) error {
	cfg := config.LoadConfig()
	
	helpText := fmt.Sprintf(`Cillers CLI version %s

Usage: cillers [command] [options]

Commands:
  new <name>        Create a new system with the specified name
  help              Show this help message
  version           Show the version number
  coder [task]      AI coding assistant that follows the instructions in .cillers/context files.
  coder-init        Initialize a new .cillers/context directory with template files
  info <request>    Get information about the project
  review            A code review with general feedback and specific violoations of the specified directives.

Options:
  --verbose         Enable verbose output for debugging purposes

Examples:
  cillers new my-project
  cillers new my-project --verbose
  cillers help
  cillers coder
  cillers coder "Rename the foo function to bar."
  cillers info "Create a list of all the functions in this project and list the functions they are used by. Are any of the functions unused in this project?"
  cillers version
  cillers review
  cillers coder-init

For more information, please visit: %s
`, cfg.Version, cfg.DocumentationURL)

	fmt.Println(helpText)
	
	return nil
}
