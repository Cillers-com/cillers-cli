package commands

import (
    "fmt"
    "os"
    "path/filepath"

    "cillers-cli/coder/templates"
    "cillers-cli/lib"
)

func CoderInit(parsedArgs lib.ParsedArgs) error {
    verbose := parsedArgs.BoolOptions["verbose"]
    currentDir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("error getting current directory: %w", err)
    }

    contextDir := filepath.Join(currentDir, ".cillers", "context")
    if err := os.MkdirAll(contextDir, 0755); err != nil {
        return fmt.Errorf("error creating context directory: %w", err)
    }

    if verbose {
        fmt.Printf("Creating .cillers/context directory in %s\n", currentDir)
    }

    files := map[string]string{
        "directives/general":       templates.GeneralDirectivesTemplate,
        "directives/language_go":   templates.LanguageDirectivesGoTemplate,
        "directives/project":       templates.ProjectDirectivesTemplate,
        "ignore":                   templates.IgnoreTemplate,
        "response_coder":           templates.ResponseCoderTemplate,
        "response_review":          templates.ResponseReviewTemplate,
        "task":                     templates.TaskTemplate,
    }

    for file, content := range files {
        filePath := filepath.Join(contextDir, file)
        dirPath := filepath.Dir(filePath)
        if err := os.MkdirAll(dirPath, 0755); err != nil {
            return fmt.Errorf("error creating directory %s: %w", dirPath, err)
        }

        if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
            return fmt.Errorf("error writing file %s: %w", filePath, err)
        }

        if verbose {
            fmt.Printf("Created file: %s\n", filePath)
        }
    }

    fmt.Println("Successfully initialized .cillers/context directory with template files.")
    return nil
}
