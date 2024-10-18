package commands

import (
    "fmt"
    "os"

    "cillers-cli/coder"
)

func Review(args []string, options map[string]bool) error {
    verbose := options["verbose"]

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

    prompt, err := coder.GenerateReviewPrompt(fileContents)
    if err != nil {
        return fmt.Errorf("error building prompt: %w", err)
    }

    if verbose {
        fmt.Println("Generated prompt for Coder:")
    }
    fmt.Print(prompt)
    return nil
}
