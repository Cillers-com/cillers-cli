package commands

import (
	"fmt"
	"cillers-cli/lib"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type Config struct {
	Stack []string `yaml:"cillers"`
}

func Start(args []string, options map[string]bool) error {
	verbose := options["verbose"]
	if verbose {
		fmt.Println("Starting Cillers ...")
	}

	configFile, err := ioutil.ReadFile("./conf/cillers.yml")
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("error parsing config file: %v", err)
	}

	if len(config.Stack) == 0 {
		return fmt.Errorf("no stack items found in the configuration")
	}

	command := "pt"
	args = append([]string{"pt", "run"}, config.Stack...)

	if verbose {
		fmt.Printf("Executing command: %s %s\n", command, strings.Join(args[1:], " "))
	}

	return lib.ExecuteTakeOverCurrentProcess(command, args)
}
