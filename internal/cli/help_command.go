package cli

import (
	"flag"
	"fmt"
)

type HelpCommand struct {
	programName string
	commands    []Command
	commandName string
	fs          *flag.FlagSet
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
	c.fs.StringVar(&c.commandName, "command", "", "Command to get help for")
}

func (c *HelpCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *HelpCommand) Usage() {
	c.fs.Usage()
}

func (c *HelpCommand) Help() {
	fmt.Println("help -command <command>")
}

func (c *HelpCommand) Execute(s *State) error {
	for _, sc := range c.commands {
		if sc.Name() == c.commandName {
			sc.Help()
			sc.Usage()
			return nil
		}
	}

	fmt.Printf("Usage: %s <subcommand> [args]\n", c.programName)
	fmt.Println("Subcommands:")

	for _, c := range c.commands {
		fmt.Printf("  ")
		c.Help()
	}

	return nil
}
