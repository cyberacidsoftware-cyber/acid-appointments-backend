package cli

import (
	"fmt"
	"sort"
)

type Cli struct {
	help *HelpCommand
}

func NewCli(programName string) *Cli {
	help := HelpCommand{programName: programName}
	help.commands = []Command{
		&LoginCommand{},
		&CreateAdminCommand{},
		&CreateSecretaryCommand{},
		&CreateCustomerCommand{},
		&CreateSpecialistCommand{},
		&CreateSpecializationCommand{},
		&CreateServiceCommand{},
		&ListAppointmentsCommand{},
		&ListSpecializationsCommand{},
		&RunInputCommand{},
		&ResetDbCommand{},
		&help,
	}

	sort.Slice(help.commands, func(i, j int) bool {
		return help.commands[i].Name() < help.commands[j].Name()
	})

	for _, c := range help.commands {
		c.Init()
	}

	return &Cli{help: &help}
}

func (c *Cli) Run(args []string) {
	if len(args) == 0 {
		c.help.Execute(nil)
		return
	}

	for _, c := range c.help.commands {
		if c.Name() == args[0] {
			s := NewState()
			defer s.Close()
			defer s.tx.Rollback(s.ctx)

			c.Parse(args[1:])
			err := c.Execute(s)

			if err != nil {
				fmt.Printf("ERROR: %s", err.Error())
				return
			}

			if err := s.tx.Commit(s.ctx); err != nil {
				fmt.Printf("ERROR: %s", err.Error())
				return
			}
			return
		}
	}

	fmt.Printf("ERROR: Unknown subcommand: %s\n", args[0])
}
