package coder

import (
    "fmt"
    "path/filepath"
    "sort"
    "strings"
    "html"
)

func escapeXML(s string) string {
    return html.EscapeString(s)
}

func PrintSortedFileContents(fileContents map[string]string) {
    sortedFiles := getSortedFiles(fileContents)
    for _, path := range sortedFiles {
        fmt.Printf("%s:\n%s\n\n", path, fileContents[path])
    }
}

func PrintSortedFileContentsToBuilder(builder *strings.Builder, fileContents map[string]string, ignoreFile string) {
    ignorePatterns, err := LoadIgnorePatterns(ignoreFile)
    if err != nil {
        fmt.Printf("Error reading ignore patterns: %v\n", err)
        ignorePatterns = []string{}
    }

    sortedFiles := getSortedFiles(fileContents)
    for _, path := range sortedFiles {
        if !matchesIgnorePattern(path, ignorePatterns) {
            builder.WriteString(fmt.Sprintf("EXISTING FILE %s:\n%s\n\n\n", path, fileContents[path]))
        }
    }
}

func getSortedFiles(fileContents map[string]string) []string {
    var currentDirFiles, otherFiles []string

    for path := range fileContents {
        if filepath.Dir(path) == "." {
            currentDirFiles = append(currentDirFiles, path)
        } else {
            otherFiles = append(otherFiles, path)
        }
    }

    sort.Strings(currentDirFiles)
    sort.Strings(otherFiles)

    return append(currentDirFiles, otherFiles...)
}

func PrintSortedFileContentsToXMLBuilder(builder *strings.Builder, fileContents map[string]string, ignoreFile string) {
    ignorePatterns, err := LoadIgnorePatterns(ignoreFile)
    if err != nil {
        fmt.Printf("Error reading ignore patterns: %v\n", err)
        ignorePatterns = []string{}
    }

    sortedFiles := getSortedFiles(fileContents)
    for _, path := range sortedFiles {
        if !matchesIgnorePattern(path, ignorePatterns) {
            builder.WriteString(fmt.Sprintf("    <file path=\"%s\">\n      <content>\n%s\n      </content>\n    </file>\n", escapeXML(path), escapeXML(fileContents[path])))
        }
    }
}
