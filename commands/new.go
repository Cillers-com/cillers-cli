package commands

import (
	"fmt"

	"cillers-cli/config"
	"cillers-cli/lib"
)

// CommandNew handles the 'new' command
func CommandNew(args []string, options map[string]bool) error {
	if len(args) == 0 {
		return fmt.Errorf("no name provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("command 'new' takes only one argument")
	}

	name := args[0]
	verbose := options["verbose"]

	if err := lib.AssertDoesntExist(name); err != nil {
		return fmt.Errorf("invalid argument: %w", err)
	}

	if !lib.IsGitInstalled() {
		return fmt.Errorf("Git is not installed or not in the PATH")
	}

	cfg := config.Get()
	fmt.Printf("Creating new system named '%s'...\n", name)

	if err := lib.Clone(cfg.TemplateRepoURL, name, verbose); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	if err := lib.Reset(name, verbose); err != nil {
		return fmt.Errorf("failed to reset repository: %w", err)
	}

	fmt.Printf("New Cillers system '%s' successfully created.\n", name)
	return nil
}
