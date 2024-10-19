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
		_ = commands.Help(nil, nil)
		os.Exit(1)
	}
}

func registerCommands() {
	lib.RegisterCommand("new", commands.New)
	lib.RegisterCommand("help", commands.Help)
	lib.RegisterCommand("version", commands.Version)
	lib.RegisterCommand("start", commands.Start)
	lib.RegisterCommand("coder", commands.Coder)
	lib.RegisterCommand("info", commands.Info)
	lib.RegisterCommand("review", commands.Review)
	lib.RegisterCommand("coder-init", commands.CoderInit)
	lib.RegisterCommand("add-commit-msg-hook", commands.AddCommitMsgHook)
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

	return fn(parsedArgs.Args, parsedArgs.BoolOptions)
}
