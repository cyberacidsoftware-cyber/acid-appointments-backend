package cli

import (
	"flag"
	"fmt"
	"log"

	"backend/internal/usecase"
)

type CreateServiceCommand struct {
	Args usecase.CreateServiceNameArgs
	fs   *flag.FlagSet
}

func (c *CreateServiceCommand) Name() string {
	return "create-service"
}

func (c *CreateServiceCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
	c.fs.StringVar(&c.Args.Name, "name", "", "Name")
	c.fs.StringVar(&c.Args.SpecializationIDRaw, "specialization-id", "", "Specialization ID")
}

func (c *CreateServiceCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *CreateServiceCommand) Usage() {
	c.fs.Usage()
}

func (c *CreateServiceCommand) Help() {
	fmt.Println("create-service -name <name>")
}

func (c *CreateServiceCommand) Execute(s *State) error {
	if err := c.Args.Validate(); err != nil {
		return err.Error
	}

	id, err := usecase.CreateServiceName(s, c.Args)
	if err != nil {
		return err.Error
	}

	log.Printf("Service created with name: %s\n%s", c.Args.Name, id)
	return nil
}
