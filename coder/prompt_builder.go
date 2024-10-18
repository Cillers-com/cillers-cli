package coder

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "html"
)

func escapeXML(s string) string {
    return html.EscapeString(s)
}

func GenerateReviewPrompt(fileContents map[string]string) (string, error) {
    promptDir := ".cillers/context"
    directivesDir := filepath.Join(promptDir, "directives")
    var task string

    var prompt strings.Builder

    prompt.WriteString("<prompt>\n")

    task = "Review the project's code to see if it fulfills each of the directives specified in the 'directive' elements in the files in the .cillers/context/directives directory. Verify that the directive is really violated before stating that it is."

    writeXMLSection(&prompt, "task_specification", "", task)

    response, err := readFileContent(filepath.Join(promptDir, "response_review"))
    if err != nil {
        return "", err
    }
    writeXMLSection(&prompt, "response_specification", "Your response should be formatted in the following way.", response)

    prompt.WriteString("  <directives>\n")
    err = filepath.Walk(directivesDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            content, err := readFileContent(path)
            if err != nil {
                return err
            }
            writeXMLSection(&prompt, fmt.Sprintf("directive_%s", strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))), "", content)
        }
        return nil
    })
    if err != nil {
        return "", fmt.Errorf("error reading directives: %w", err)
    }
    prompt.WriteString("  </directives>\n")

    prompt.WriteString("  <existing_project_files>\n")
    PrintSortedFileContentsToXMLBuilder(&prompt, fileContents, filepath.Join(promptDir, "ignore"))
    prompt.WriteString("  </existing_project_files>\n")

    prompt.WriteString("</prompt>")

    return prompt.String(), nil
}
func GenerateCoderPrompt(fileContents map[string]string, task string) (string, error) {
    promptDir := ".cillers/context"
    directivesDir := filepath.Join(promptDir, "directives")

    var prompt strings.Builder

    prompt.WriteString("<prompt>\n")

    if task == "" {
        taskSpec, err := readFileContent(filepath.Join(promptDir, "task"))
        if err != nil {
            return "", err
        }
        writeXMLSection(&prompt, "task_specification", "Accomplish the following tasks by following the directives, facts and restrictions provided below.", taskSpec)
    } else {
        writeXMLSection(&prompt, "task_specification", "Accomplish the following tasks by following the directives, facts and restrictions provided below.", task)
    }

    response, err := readFileContent(filepath.Join(promptDir, "response_coder"))
    if err != nil {
        return "", err
    }
    writeXMLSection(&prompt, "response_specification", "Your response should be formatted in the following way.", response)

    prompt.WriteString("  <directives>\n")
    prompt.WriteString("    <header>Any response that fails to comply with the directives will be rejected. Please ensure strict adherence to these directives.</header>\n")
    err = filepath.Walk(directivesDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            content, err := readFileContent(path)
            if err != nil {
                return err
            }
            writeXMLSection(&prompt, fmt.Sprintf("directive_%s", strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))), "", content)
        }
        return nil
    })
    if err != nil {
        return "", fmt.Errorf("error reading directives: %w", err)
    }
    prompt.WriteString("  </directives>\n")

    prompt.WriteString("  <existing_project_files>\n")
    PrintSortedFileContentsToXMLBuilder(&prompt, fileContents, filepath.Join(promptDir, "ignore"))
    prompt.WriteString("  </existing_project_files>\n")

    prompt.WriteString("</prompt>")

    return prompt.String(), nil
}

func BuildInfoPrompt(request string, fileContents map[string]string) (string, error) {
    var prompt strings.Builder

    prompt.WriteString("<prompt>\n")

    // Request
    writeXMLSection(&prompt, "request", "", request)

    // Existing Project Files
    prompt.WriteString("  <existing_project_files>\n")
    PrintSortedFileContentsToXMLBuilder(&prompt, fileContents, ".cillers/context/ignore")
    prompt.WriteString("  </existing_project_files>\n")

    prompt.WriteString("</prompt>")

    return prompt.String(), nil
}

func readFileContent(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("error reading file %s: %w", path, err)
    }
    return string(content), nil
}

func writeXMLSection(builder *strings.Builder, tag, header, content string) {
    builder.WriteString(fmt.Sprintf("  <%s>\n", tag))
    if header != "" {
        builder.WriteString(fmt.Sprintf("    <header>%s</header>\n", header))
    }
    builder.WriteString(fmt.Sprintf("    <content>%s</content>\n", content))
    builder.WriteString(fmt.Sprintf("  </%s>\n", tag))
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
