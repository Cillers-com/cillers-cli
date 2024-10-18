package coder

import (
    "path/filepath"
    "strings"
)

func matchesIgnorePattern(path string, ignorePatterns []string) bool {
    for _, pattern := range ignorePatterns {
        if globMatch(pattern, path) {
            return true
        }
    }
    return false
}

func globMatch(pattern, path string) bool {
    if !strings.Contains(pattern, "**") {
        matched, _ := filepath.Match(pattern, path)
        return matched
    }

    parts := strings.Split(pattern, "**")
    if !strings.HasPrefix(path, parts[0]) {
        return false
    }

    path = path[len(parts[0]):]
    for i := 1; i < len(parts)-1; i++ {
        idx := strings.Index(path, parts[i])
        if idx == -1 {
            return false
        }
        path = path[idx+len(parts[i]):]
    }

    return strings.HasSuffix(path, parts[len(parts)-1])
}
