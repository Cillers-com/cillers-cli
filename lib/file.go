package lib 

import (
	"fmt"
	"os"
)

// Exists checks if a file or directory exists
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// AssertDoesntExist checks that a file or directory doesn't exist and returns an error if it does
func AssertDoesntExist(name string) error {
	if Exists(name) {
		return fmt.Errorf("file or directory already exists: %s", name)
	}
	return nil
}

// AssertExists checks that a file or directory exists and returns an error if it doesn't
func AssertExists(name string) error {
	if !Exists(name) {
		return fmt.Errorf("file or directory does not exist: %s", name)
	}
	return nil
}
