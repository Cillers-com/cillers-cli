package commands

import (
    "fmt"
    "os"
    "strings"

    "cillers-cli/coder"
)

func Info(args []string, options map[string]bool) error {
    if len(args) == 0 {
        return fmt.Errorf("no request provided")
    }

    request := strings.Join(args, " ")
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

    prompt, err := coder.BuildInfoPrompt(request, fileContents)
    if err != nil {
        return fmt.Errorf("error building prompt: %w", err)
    }

    if verbose {
        fmt.Println("Generated prompt for Info:")
    }
    fmt.Print(prompt)
    return nil
}
