package commands

import (
	"fmt"

	"cillers-cli/config"
)

// CommandVersion handles the 'version' command
func CommandVersion(args []string, options map[string]bool) error {
	cfg := config.Get()
	fmt.Printf("Cillers CLI version %s\n", cfg.Version)
	return nil
}
