package cli

import (
	"flag"
	"fmt"
	"log"

	"backend/internal/usecase"
)

type CreateSpecializationCommand struct {
	Args usecase.SpecializationInfoArgs
	fs   *flag.FlagSet
}

func (c *CreateSpecializationCommand) Name() string {
	return "create-specialization"
}

func (c *CreateSpecializationCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
	c.fs.StringVar(&c.Args.Name, "name", "", "Name")
}

func (c *CreateSpecializationCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *CreateSpecializationCommand) Usage() {
	c.fs.Usage()
}

func (c *CreateSpecializationCommand) Help() {
	fmt.Println("create-specialization -name <name>")
}

func (c *CreateSpecializationCommand) Execute(s *State) error {
	if err := c.Args.Validate(); err != nil {
		return err.Error
	}

	id, err := usecase.CreateSpecialization(s, c.Args)
	if err != nil {
		return err.Error
	}

	log.Printf("Specialization created with name: %s\n%s", c.Args.Name, id)
	return nil
}
