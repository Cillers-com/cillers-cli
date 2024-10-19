package commands

import (
    "fmt"

    "cillers-cli/config"
    "cillers-cli/lib"
)

func Version(parsedArgs lib.ParsedArgs) error {
    cfg := config.LoadConfig()
    fmt.Printf("Cillers CLI version %s\n", cfg.Version)
    return nil
}
