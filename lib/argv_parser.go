package lib

import (
	"fmt"
	"strings"
)

// ParsedArgs represents the parsed command-line arguments
type ParsedArgs struct {
	Command string
	Args    []string
	Options map[string]bool
}

// ParseArgv parses the command-line arguments
func ParseArgv(args []string) (ParsedArgs, error) {
	// Handle help and version as special cases
	if contains(args, "--help") {
		return ParsedArgs{Command: "help", Args: []string{}, Options: map[string]bool{}}, nil
	}
	if contains(args, "--version") {
		return ParsedArgs{Command: "version", Args: []string{}, Options: map[string]bool{}}, nil
	}

	idx := indexOfFirstOption(args)
	commandArgs := args[:idx]
	optionArgs := args[idx:]

	if len(commandArgs) == 0 {
		return ParsedArgs{}, fmt.Errorf("no command provided")
	}

	command := commandArgs[0]
	if !IsSupportedCommand(command) {
		return ParsedArgs{}, fmt.Errorf("unsupported command: %s", command)
	}

	options, err := parseOptions(optionArgs)
	if err != nil {
		return ParsedArgs{}, err
	}

	return ParsedArgs{
		Command: command,
		Args:    commandArgs[1:],
		Options: options,
	}, nil
}

func indexOfFirstOption(args []string) int {
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			return i
		}
	}
	return len(args)
}

func parseOptions(args []string) (map[string]bool, error) {
	options := make(map[string]bool)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			option := strings.TrimPrefix(arg, "--")
			if !IsSupportedOption(option) {
				return nil, fmt.Errorf("unsupported option: %s", arg)
			}
			options[option] = true
		}
	}
	return options, nil
}

// IsOptionSet checks if a specific option was set in the command line arguments
func (pa ParsedArgs) IsOptionSet(option string) bool {
	return pa.Options[option]
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
