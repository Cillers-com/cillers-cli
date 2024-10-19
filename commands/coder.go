package commands

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "cillers-cli/coder"
    "cillers-cli/lib"
)

func Coder(args []string, options map[string]bool) error {
    verbose := options["verbose"]
    var task string

    isClean, err := lib.IsWorkingTreeClean()
    if err != nil {
        return fmt.Errorf("error checking Git working tree: %w", err)
    }

    if !isClean {
        fmt.Println("Warning: The Git working tree is not clean.")
        fmt.Print("Do you want to proceed anyway? (y/N): ")
        reader := bufio.NewReader(os.Stdin)
        response, err := reader.ReadString('\n')
        if err != nil {
            return fmt.Errorf("error reading user input: %w", err)
        }
        response = strings.TrimSpace(strings.ToLower(response))
        if response != "y" && response != "yes" {
            fmt.Println("Aborting operation.")
            return nil
        }
    }

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
