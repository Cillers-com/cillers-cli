package coder

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func LoadIgnorePatterns(ignoreFile string) ([]string, error) {
    patterns := []string{ignoreFile}

    file, err := os.Open(ignoreFile)
    if err != nil {
        return nil, fmt.Errorf("error opening ignore file %s: %w", ignoreFile, err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        patterns = append(patterns, strings.TrimSpace(scanner.Text()))
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading ignore file %s: %w", ignoreFile, err)
    }

    return patterns, nil
}
