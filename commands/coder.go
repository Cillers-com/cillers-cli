package commands

import (
    "fmt"
    "os"
    "strings"

    "cillers-cli/coder"
)

func Coder(args []string, options map[string]bool) error {
    verbose := options["verbose"]
    var task string

    if len(args) > 0 {
        task = strings.Join(args, " ")
    }

    ignorePatterns, err := coder.LoadIgnorePatterns(".cillers/context/ignore")
    if err != nil {
        return fmt.Errorf("error reading ignore patterns: %w", err)
    }

    currentDir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("error getting current directory: %w", err)
    }

    fileContents, err := coder.LoadFileContents(currentDir, ignorePatterns)
    if err != nil {
        return fmt.Errorf("error walking the directory tree: %w", err)
    }

    prompt, err := coder.GenerateCoderPrompt(fileContents, task)
    if err != nil {
        return fmt.Errorf("error building prompt: %w", err)
    }

    if verbose {
        fmt.Println("Generated prompt for Coder:")
    }
    fmt.Print(prompt)
    return nil
}
