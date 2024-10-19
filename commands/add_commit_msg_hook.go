package commands

import (
    "fmt"
    "os"
    "path/filepath"
    "cillers-cli/coder/templates"
    "cillers-cli/lib"
)

func AddCommitMsgHook(parsedArgs lib.ParsedArgs) error {
    if !isGitRepository() {
        return fmt.Errorf("not a git repository")
    }

    hookPath := filepath.Join(".git", "hooks", "prepare-commit-msg")
    hookContent := templates.GenerateCommitMsgHook()

    err := os.WriteFile(hookPath, []byte(hookContent), 0755)
    if err != nil {
        return fmt.Errorf("failed to create prepare-commit-msg hook: %w", err)
    }

    fmt.Println("prepare-commit-msg hook added successfully.")
    return nil
}

func isGitRepository() bool {
    _, err := os.Stat(".git")
    return !os.IsNotExist(err)
}
