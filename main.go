package main

import (
	"fmt"
	"os"

	"cillers-cli/commands"
	"cillers-cli/lib"
)

func main() {
	registerCommands()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		_ = commands.CommandHelp(nil, nil)
		os.Exit(1)
	}
}

func registerCommands() {
	lib.RegisterCommand("new", commands.CommandNew)
	lib.RegisterCommand("help", commands.CommandHelp)
	lib.RegisterCommand("version", commands.CommandVersion)
}

func run() error {
	parsedArgs, err := lib.ParseArgv(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse arguments: %w", err)
	}

	fn, ok := lib.GetCommand(parsedArgs.Command)
	if !ok {
		return fmt.Errorf("unknown command: %s", parsedArgs.Command)
	}

	return fn(parsedArgs.Args, parsedArgs.Options)
}
