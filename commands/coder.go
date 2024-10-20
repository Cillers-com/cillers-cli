package commands

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "path/filepath"

    "cillers-cli/coder"
    "cillers-cli/lib"
)

func Coder(parsedArgs lib.ParsedArgs) error {
    verbose := parsedArgs.BoolOptions["verbose"]
    var task string

    isClean, err := lib.IsWorkingTreeClean(".cillers/context/task")
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

    if len(parsedArgs.Args) > 0 {
        task = strings.Join(parsedArgs.Args, " ")
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
        fmt.Print(prompt)
    }

    changeProposal, err := lib.SendPromptToAnthropic(prompt)
    if err != nil {
        return fmt.Errorf("error sending prompt to Anthropic: %w", err)
    }

    fmt.Println("\nParsed Change Proposal:")
    fmt.Printf("Description:\n")
    fmt.Printf("  Change Summary: %s\n", changeProposal.Description.ChangeSummary)
    for _, detail := range changeProposal.Description.ChangeDetails {
        fmt.Printf("  Change Detail: %s\n", detail)
    }

    fmt.Printf("\nSpecification:\n")
    fmt.Printf("  Files to be created:\n")
    for _, file := range changeProposal.Specification.FilesToBeCreated {
        fmt.Printf("    Path: %s\n", file.Path)
        fmt.Printf("    Content:\n%s\n", file.Content)
    }
    fmt.Printf("  Files to be updated:\n")
    for _, file := range changeProposal.Specification.FilesToBeUpdated {
        fmt.Printf("    Path: %s\n", file.Path)
        fmt.Printf("    Content:\n%s\n", file.Content)
    }
    fmt.Printf("  Files to be deleted:\n")
    for _, file := range changeProposal.Specification.FilesToBeDeleted {
        fmt.Printf("    Path: %s\n", file.Path)
    }

    fmt.Printf("\nCode Review:\n")
    fmt.Printf("  Positive Feedback: %s\n", changeProposal.CodeReview.PositiveFeedback)
    fmt.Printf("  Improvement Suggestions: %s\n", changeProposal.CodeReview.ImprovementSuggestions)
    fmt.Printf("  Code Quality Assessment: %s\n", changeProposal.CodeReview.CodeQualityAssessment)

    // Confirmation before applying changes
    fmt.Print("\nDo you want to apply these changes? (y/N): ")
    reader := bufio.NewReader(os.Stdin)
    response, err := reader.ReadString('\n')
    if err != nil {
        return fmt.Errorf("error reading user input: %w", err)
    }
    response = strings.TrimSpace(strings.ToLower(response))
    if response != "y" && response != "yes" {
        fmt.Println("Changes not applied.")
        return nil
    }

    // Apply changes
    for _, file := range changeProposal.Specification.FilesToBeCreated {
        err := os.MkdirAll(filepath.Dir(file.Path), 0755)
        if err != nil {
            return fmt.Errorf("error creating directory for file %s: %w", file.Path, err)
        }
        err = os.WriteFile(file.Path, []byte(file.Content), 0644)
        if err != nil {
            return fmt.Errorf("error creating file %s: %w", file.Path, err)
        }
        fmt.Printf("Created file: %s\n", file.Path)
    }

    for _, file := range changeProposal.Specification.FilesToBeUpdated {
        err := os.WriteFile(file.Path, []byte(file.Content), 0644)
        if err != nil {
            return fmt.Errorf("error updating file %s: %w", file.Path, err)
        }
        fmt.Printf("Updated file: %s\n", file.Path)
    }

    for _, file := range changeProposal.Specification.FilesToBeDeleted {
        err := os.Remove(file.Path)
        if err != nil {
            return fmt.Errorf("error deleting file %s: %w", file.Path, err)
        }
        fmt.Printf("Deleted file: %s\n", file.Path)
    }

    fmt.Println("Changes applied successfully.")
    return nil
}
