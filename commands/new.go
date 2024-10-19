package commands

import (
    "fmt"
    "cillers-cli/config"
    "cillers-cli/lib"
)

func New(parsedArgs lib.ParsedArgs) error {
    if len(parsedArgs.Args) == 0 {
        return fmt.Errorf("no name provided")
    }
    if len(parsedArgs.Args) > 2 {
        return fmt.Errorf("command 'new' takes at most two arguments")
    }

    name := parsedArgs.Args[0]
    branch := "main" // Default branch

    if len(parsedArgs.Args) == 2 {
        branch = parsedArgs.Args[1]
    }

    verbose := parsedArgs.BoolOptions["verbose"]

    if err := lib.AssertDoesntExist(name); err != nil {
        return fmt.Errorf("invalid argument: %w", err)
    }

    if !lib.IsGitInstalled() {
        return fmt.Errorf("Git is not installed or not in the PATH")
    }

    cfg := config.LoadConfig()
    fmt.Printf("Creating new system named '%s' from branch '%s'...\n", name, branch)

    if err := lib.Clone(cfg.TemplateRepoURL, name, branch, verbose); err != nil {
        return fmt.Errorf("failed to clone repository: %w", err)
    }

    if err := lib.Reset(name, verbose); err != nil {
        return fmt.Errorf("failed to reset repository: %w", err)
    }

    fmt.Printf("New Cillers system '%s' successfully created from branch '%s'.\n", name, branch)
    return nil
}
