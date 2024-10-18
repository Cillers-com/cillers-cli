package lib

type CommandFunc func(args []string, options map[string]bool) error

var commandRegistry = make(map[string]CommandFunc)

func RegisterCommand(name string, fn CommandFunc) {
	commandRegistry[name] = fn
}

func GetCommand(name string) (CommandFunc, bool) {
	fn, ok := commandRegistry[name]
	return fn, ok
}

func IsSupportedCommand(name string) bool {
	_, ok := commandRegistry[name]
	return ok
}

func IsSupportedOption(option string) bool {
	// Add your supported options here
	supportedOptions := map[string]bool{
		"verbose": true,
		// Add more options as needed
	}
	return supportedOptions[option]
}

func GetSupportedCommands() []string {
	commands := make([]string, 0, len(commandRegistry))
	for cmd := range commandRegistry {
		commands = append(commands, cmd)
	}
	return commands
}

