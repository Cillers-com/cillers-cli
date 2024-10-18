package commands

import (
	"fmt"

	"cillers-cli/config"
)

func Version(args []string, options map[string]bool) error {
	cfg := config.LoadConfig()
	fmt.Printf("Cillers CLI version %s\n", cfg.Version)
	return nil
}
