package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
    "syscall"
)

// ExecResult represents the result of a command execution
type ExecResult struct {
	Stdout string
	Stderr string
}

// Execute runs a command in the specified directory with real-time output
func Execute(path string, commandArgs []string, verbose bool) (ExecResult, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return ExecResult{}, fmt.Errorf("failed to get absolute path: %w", err)
	}

	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)
	cmd.Dir = absPath

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return ExecResult{}, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return ExecResult{}, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if verbose {
		fmt.Printf("Executing in %s: %s\n", absPath, strings.Join(commandArgs, " "))
	}

	err = cmd.Start()
	if err != nil {
		return ExecResult{}, fmt.Errorf("failed to start command: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var stdoutBuilder, stderrBuilder strings.Builder

	go func() {
		defer wg.Done()
		scanAndWrite(stdout, &stdoutBuilder, os.Stdout, "stdout: ", verbose)
	}()

	go func() {
		defer wg.Done()
		scanAndWrite(stderr, &stderrBuilder, os.Stderr, "stderr: ", verbose)
	}()

	wg.Wait()

	err = cmd.Wait()

	result := ExecResult{
		Stdout: stdoutBuilder.String(),
		Stderr: stderrBuilder.String(),
	}

	if err != nil {
		return result, fmt.Errorf("command execution failed: %w", err)
	}

	return result, nil
}

func scanAndWrite(r io.Reader, builder *strings.Builder, w io.Writer, prefix string, verbose bool) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if verbose {
			fmt.Fprintln(w, prefix+line)
		}
		builder.WriteString(line + "\n")
	}
}

// ExecuteVerbose is a wrapper around Execute that always prints output
func ExecuteVerbose(path string, commandArgs []string) (ExecResult, error) {
	return Execute(path, commandArgs, true)
}

// Exec replaces the current process with a new command
func ExecuteTakeOverCurrentProcess(command string, args []string) error {
	path, err := exec.LookPath(command)
	if err != nil {
		return fmt.Errorf("command '%s' not found: %w", command, err)
	}

	if err := syscall.Exec(path, args, os.Environ()); err != nil {
		return fmt.Errorf("failed to execute '%s': %w", command, err)
	}

	// This line will never be reached if Exec is successful
	return nil
}
