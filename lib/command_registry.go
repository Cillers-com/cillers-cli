package lib

type CommandFunc func(parsedArgs ParsedArgs) error

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
    supportedOptions := map[string]bool{
        "verbose": true,
        "help":    true,
        "version": true,
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
