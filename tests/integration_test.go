package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var (
	projectRoot string
	tmpTestDir  string
)

func TestMain(m *testing.M) {
	var err error
	projectRoot, err = filepath.Abs("../")
	if err != nil {
		panic("Failed to get project root: " + err.Error())
	}

	// Create a temporary directory for all tests
	tmpTestDir = filepath.Join(projectRoot, "tests", "tmp")
	err = os.MkdirAll(tmpTestDir, 0755)
	if err != nil {
		panic("Failed to create temp test directory: " + err.Error())
	}

	// Run tests
	code := m.Run()

	// Clean up
	os.RemoveAll(tmpTestDir)

	os.Exit(code)
}

func TestNewCommand(t *testing.T) {
	testDir := filepath.Join(tmpTestDir, "new-command-test")
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)
	
	err = os.Chdir(testDir)
	if err != nil {
		t.Fatalf("Failed to change to test directory: %v", err)
	}

	cmd := exec.Command("go", "run", filepath.Join(projectRoot, "main.go"), "new", "test-project")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	expectedOutput := "New Cillers system 'test-project' successfully created from branch 'main'."
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Expected output to contain '%s', but got: %s", expectedOutput, output)
	}

	if _, err := os.Stat("test-project"); os.IsNotExist(err) {
		t.Errorf("Project directory was not created")
	}
}

func TestHelpCommand(t *testing.T) {
	cmd := exec.Command("go", "run", filepath.Join(projectRoot, "main.go"), "help")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	expectedStrings := []string{
		"Cillers CLI version",
		"Usage: cillers [command] [options]",
		"Commands:",
		"new <name>",
		"help",
		"version",
		"coder [task]",
		"info <request>",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(string(output), expected) {
			t.Errorf("Expected output to contain '%s', but it didn't.\nOutput: %s", expected, output)
		}
	}
}

func TestVersionCommand(t *testing.T) {
	cmd := exec.Command("go", "run", filepath.Join(projectRoot, "main.go"), "version")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	expectedPrefix := "Cillers CLI version"
	if !strings.HasPrefix(string(output), expectedPrefix) {
		t.Errorf("Expected output to start with '%s', but got: %s", expectedPrefix, output)
	}
}
