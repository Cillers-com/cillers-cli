package coder

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func ReadPromptFiles() (string, error) {
    promptDir := ".cillers/coder/prompt"
    var promptContents strings.Builder

    err := filepath.Walk(promptDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error accessing path %s: %w", path, err)
        }
        if !info.IsDir() {
            content, err := os.ReadFile(path)
            if err != nil {
                return fmt.Errorf("error reading file %s: %w", path, err)
            }
            promptContents.WriteString(fmt.Sprintf("# %s\n%s\n\n", filepath.Base(path), string(content)))
        }
        return nil
    })

    if err != nil {
        return "", fmt.Errorf("error reading prompt files: %w", err)
    }

    return promptContents.String(), nil
}
