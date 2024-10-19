package lib

import (
    "fmt"
    "os/exec"
    "runtime"
    "strings"
)

func CopyToClipboard(text string) error {
    var cmd *exec.Cmd

    switch runtime.GOOS {
    case "darwin":
        cmd = exec.Command("pbcopy")
    case "windows":
        cmd = exec.Command("clip")
    default: // Linux and other Unix-like systems
        cmd = exec.Command("xclip", "-selection", "clipboard")
    }

    if cmd == nil {
        return fmt.Errorf("unsupported platform")
    }

    cmd.Stdin = strings.NewReader(text)
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("failed to copy to clipboard: %w", err)
    }

    return nil
}
