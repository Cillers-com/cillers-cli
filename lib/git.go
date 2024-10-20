package lib

import (
    "fmt"
    "strings"
    "path/filepath"
)

func IsGitInstalled() bool {
    _, err := Execute(".", []string{"git", "--version"}, false)
    return err == nil
}

func AssertOriginURL(path, expectedURL string) error {
    result, err := Execute(path, []string{"git", "remote", "get-url", "origin"}, false)
    if err != nil {
        return fmt.Errorf("failed to get origin URL: %w", err)
    }
    actualURL := strings.TrimSpace(result.Stdout)
    expectedURL = strings.TrimSpace(expectedURL)
    // Remove .git suffix if present for both URLs
    actualURL = strings.TrimSuffix(actualURL, ".git")
    expectedURL = strings.TrimSuffix(expectedURL, ".git")
    if actualURL != expectedURL {
        return fmt.Errorf("origin URL does not match expected URL.\nExpected: %s\nActual: %s", expectedURL, actualURL)
    }
    return nil
}

func Clone(url, targetName, branch string, verbose bool) error {
    args := []string{"git", "clone"}
    
    if branch != "" {
        args = append(args, "-b", branch)
    }
    
    args = append(args, url, targetName)

    _, err := Execute(".", args, verbose)
    if err != nil {
        return fmt.Errorf("failed to clone repository: %w", err)
    }
    return nil
}

func Reset(path string, verbose bool) error {
    commands := [][]string{
        {"rm", "-rf", ".git"},
        {"git", "init"},
        {"git", "add", "."},
        {"git", "commit", "-m", "Initial commit"},
    }
    for _, cmd := range commands {
        _, err := Execute(path, cmd, verbose)
        if err != nil {
            return fmt.Errorf("failed to execute command '%v': %w", cmd, err)
        }
    }
    return nil
}

func IsWorkingTreeClean(ignorePatterns ...string) (bool, error) {
    result, err := Execute(".", []string{"git", "diff", "--name-only"}, false)
    if err != nil {
        return false, fmt.Errorf("failed to check Git status: %w", err)
    }
    
    files := strings.Split(strings.TrimSpace(result.Stdout), "\n")
    for _, file := range files {
        if file == "" {
            continue
        }
        
        ignored := false
        for _, pattern := range ignorePatterns {
            matched, err := filepath.Match(pattern, file)
            if err != nil {
                return false, fmt.Errorf("invalid ignore pattern '%s': %w", pattern, err)
            }
            if matched {
                ignored = true
                break
            }
        }
        if !ignored {
            return false, nil
        }
    }
    
    return true, nil
}

