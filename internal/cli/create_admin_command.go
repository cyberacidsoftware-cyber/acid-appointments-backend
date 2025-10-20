package cli

import (
	"flag"
	"fmt"
	"log"

	"backend/internal/usecase"
)

type CreateAdminCommand struct {
	Args usecase.CreateAdminArgs
	fs   *flag.FlagSet
}

func (c *CreateAdminCommand) Name() string {
	return "create-admin"
}

func (c *CreateAdminCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
	c.fs.StringVar(&c.Args.Name, "name", "", "Name")
	c.fs.StringVar(&c.Args.Email, "email", "", "Email address")
	c.fs.StringVar(&c.Args.Password, "password", "", "Password")
}

func (c *CreateAdminCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *CreateAdminCommand) Usage() {
	c.fs.Usage()
}

func (c *CreateAdminCommand) Help() {
	fmt.Println("create-admin -name <name> -email <email> -password <password>")
}

func (c *CreateAdminCommand) Execute(s *State) error {
	if err := c.Args.Validate(); err != nil {
		return err.Error
	}

	id, err := usecase.CreateAdmin(s, c.Args)
	if err != nil {
		return fmt.Errorf("error creating admin: %v", err)
	}

	log.Printf("Admin created with email: %s\n%s", c.Args.Email, id)

	return nil
}
