package lib

import (
	"fmt"
	"strings"
)

type ParsedArgs struct {
	Command      string
	Args         []string
	BoolOptions  map[string]bool
	ValueOptions map[string]string
}

func ParseArgv(args []string) (ParsedArgs, error) {
	if contains(args, "--help") {
		return ParsedArgs{Command: "help", Args: []string{}, BoolOptions: map[string]bool{}, ValueOptions: map[string]string{}}, nil
	}
	if contains(args, "--version") {
		return ParsedArgs{Command: "version", Args: []string{}, BoolOptions: map[string]bool{}, ValueOptions: map[string]string{}}, nil
	}

	var commandArgs []string
	var optionArgs []string
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			optionArgs = args[i:]
			break
		}
		commandArgs = append(commandArgs, arg)
	}

	if len(commandArgs) == 0 {
		return ParsedArgs{}, fmt.Errorf("no command provided")
	}

	command := commandArgs[0]
	if !IsSupportedCommand(command) {
		return ParsedArgs{}, fmt.Errorf("unsupported command: %s", command)
	}

	boolOptions, valueOptions, err := parseOptions(optionArgs)
	if err != nil {
		return ParsedArgs{}, err
	}

	return ParsedArgs{
		Command:      command,
		Args:         commandArgs[1:],
		BoolOptions:  boolOptions,
		ValueOptions: valueOptions,
	}, nil
}

func parseOptions(args []string) (map[string]bool, map[string]string, error) {
	boolOptions := make(map[string]bool)
	valueOptions := make(map[string]string)
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			option := strings.TrimPrefix(arg, "--")
			if !IsSupportedOption(option) {
				return nil, nil, fmt.Errorf("unsupported option: %s", arg)
			}
			if OptionTakesValue(option) {
				if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
					valueOptions[option] = args[i+1]
					i++
				} else {
					return nil, nil, fmt.Errorf("option %s requires a value", option)
				}
			} else {
				if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
					return nil, nil, fmt.Errorf("boolean option %s should not have a value", option)
				}
				boolOptions[option] = true
			}
		}
	}
	return boolOptions, valueOptions, nil
}

func (pa ParsedArgs) IsOptionSet(option string) bool {
	if OptionTakesValue(option) {
		_, exists := pa.ValueOptions[option]
		return exists
	}
	return pa.BoolOptions[option]
}

func (pa ParsedArgs) GetOptionValue(option string) (string, bool) {
	value, exists := pa.ValueOptions[option]
	return value, exists
}

func OptionTakesValue(option string) bool {
	// Add options that should take values here
	return false
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
