package lib

import (
	"fmt"
	"strings"
)

// IsGitInstalled checks if Git is installed and available in the PATH
func IsGitInstalled() bool {
	_, err := Execute(".", []string{"git", "--version"}, false)
	return err == nil
}

// AssertOriginURL checks if the origin URL of the repository matches the expected URL
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

// Clone clones a Git repository
func Clone(url, targetName string, verbose bool) error {
	_, err := Execute(".", []string{"git", "clone", url, targetName}, verbose)
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}
	return nil
}

// Reset resets a Git repository to its initial state
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
