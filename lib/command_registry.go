package lib

// CommandFunc is the type for command handler functions
type CommandFunc func(args []string, options map[string]bool) error

// commandRegistry stores all available commands
var commandRegistry = make(map[string]CommandFunc)

// RegisterCommand adds a command to the registry
func RegisterCommand(name string, fn CommandFunc) {
	commandRegistry[name] = fn
}

// GetCommand retrieves a command function from the registry
func GetCommand(name string) (CommandFunc, bool) {
	fn, ok := commandRegistry[name]
	return fn, ok
}

// IsSupportedCommand checks if a command is supported
func IsSupportedCommand(name string) bool {
	_, ok := commandRegistry[name]
	return ok
}

// IsSupportedOption checks if an option is supported
func IsSupportedOption(option string) bool {
	// Add your supported options here
	supportedOptions := map[string]bool{
		"verbose": true,
		// Add more options as needed
	}
	return supportedOptions[option]
}

// GetSupportedCommands returns a list of all supported commands
func GetSupportedCommands() []string {
	commands := make([]string, 0, len(commandRegistry))
	for cmd := range commandRegistry {
		commands = append(commands, cmd)
	}
	return commands
}

