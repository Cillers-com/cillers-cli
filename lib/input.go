package lib

import (
    "fmt"
    "syscall"

    "golang.org/x/term"
)

func ReadPassword() (string, error) {
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        return "", fmt.Errorf("error reading password: %w", err)
    }
    fmt.Println() // Print a newline after the user presses Enter
    return string(bytePassword), nil
}
