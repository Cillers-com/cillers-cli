package coder

import (
    "fmt"
    "os"
    "path/filepath"
)

func LoadFileContents(currentDir string, ignorePatterns []string) (map[string]string, error) {
    fileContents := make(map[string]string)

    err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error accessing path %s: %w", path, err)
        }

        relPath, err := filepath.Rel(currentDir, path)
        if err != nil {
            return fmt.Errorf("error getting relative path for %s: %w", path, err)
        }

        if matchesIgnorePattern(relPath, ignorePatterns) {
            if info.IsDir() {
                return filepath.SkipDir
            }
            return nil
        }

        if !info.IsDir() {
            content, err := os.ReadFile(path)
            if err != nil {
                return fmt.Errorf("error reading file %s: %w", path, err)
            }
            fileContents[relPath] = string(content)
        }

        return nil
    })

    if err != nil {
        return nil, fmt.Errorf("error walking directory: %w", err)
    }

    return fileContents, nil
}
