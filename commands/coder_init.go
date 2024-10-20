package commands

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "cillers-cli/coder/templates"
    "cillers-cli/lib"
)

func CoderInit(parsedArgs lib.ParsedArgs) error {
    verbose := parsedArgs.BoolOptions["verbose"]
    force := parsedArgs.BoolOptions["force"]
    currentDir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("error getting current directory: %w", err)
    }

    contextDir := filepath.Join(currentDir, ".cillers", "context")
    if err := createDirIfNotExists(contextDir, verbose); err != nil {
        return err
    }

    secretsDir := filepath.Join(currentDir, ".cillers", "secrets_and_local_config")
    if err := createDirIfNotExists(secretsDir, verbose); err != nil {
        return err
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
        if err := createFileWithContent(filePath, content, verbose, force); err != nil {
            return err
        }
    }

    secretsFilePath := filepath.Join(secretsDir, "secrets.yml")
    apiKey, err := promptForAPIKey()
    if err != nil {
        return err
    }
    secretsContent := fmt.Sprintf(templates.SecretsTemplate, apiKey)
    if err := createFileWithContent(secretsFilePath, secretsContent, verbose, force); err != nil {
        return err
    }

    fmt.Println("Successfully initialized .cillers/context and .cillers/secrets_and_local_config directories with template files.")
    return nil
}

func createDirIfNotExists(path string, verbose bool) error {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        if err := os.MkdirAll(path, 0755); err != nil {
            return fmt.Errorf("error creating directory %s: %w", path, err)
        }
        if verbose {
            fmt.Printf("Created directory: %s\n", path)
        }
    } else if verbose {
        fmt.Printf("Directory already exists: %s\n", path)
    }
    return nil
}

func createFileWithContent(path string, content string, verbose bool, force bool) error {
    dirPath := filepath.Dir(path)
    if err := createDirIfNotExists(dirPath, verbose); err != nil {
        return err
    }

    if _, err := os.Stat(path); err == nil && !force {
        if !confirmOverwrite(path) {
            fmt.Printf("Skipping file: %s\n", path)
            return nil
        }
    }

    if err := os.WriteFile(path, []byte(content), 0644); err != nil {
        return fmt.Errorf("error writing file %s: %w", path, err)
    }

    if verbose {
        fmt.Printf("Created file: %s\n", path)
    }
    return nil
}

func confirmOverwrite(path string) bool {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("File %s already exists. Overwrite? (y/n): ", path)
        response, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            return false
        }
        response = strings.ToLower(strings.TrimSpace(response))
        if response == "y" || response == "yes" {
            return true
        } else if response == "n" || response == "no" {
            return false
        }
        fmt.Println("Please answer with 'y' or 'n'.")
    }
}

func promptForAPIKey() (string, error) {
    fmt.Print("Enter your Anthropic API key: ")
    apiKey, err := lib.ReadPassword()
    if err != nil {
        return "", fmt.Errorf("error reading API key: %w", err)
    }
    return apiKey, nil
}
