package commands

import (
    "fmt"

    "cillers-cli/config"
    "cillers-cli/lib"
)

func Help(parsedArgs lib.ParsedArgs) error {
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
  review            A code review with general feedback and specific violations of the specified directives.
  add-commit-msg-hook Add a Git prepare-commit-msg hook that uses Claude Sonnet API to generate commit messages

Options:
  --verbose         Enable verbose output for debugging purposes
  --force           Skip user confirmation prompts and apply changes automatically

Examples:
  cillers new my-project
  cillers new my-project --verbose
  cillers help
  cillers coder
  cillers coder "Rename the foo function to bar."
  cillers coder --force "Rename the foo function to bar."
  cillers info "Create a list of all the functions in this project and list the functions they are used by. Are any of the functions unused in this project?"
  cillers version
  cillers review
  cillers coder-init
  cillers coder-init --force
  cillers add-commit-msg-hook

For more information, please visit: %s
`, cfg.Version, cfg.DocumentationURL)

    fmt.Println(helpText)
    
    return nil
}