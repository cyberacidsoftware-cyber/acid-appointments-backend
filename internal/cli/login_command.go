package cli

import (
	"flag"
	"fmt"
	"log"
)

// TODO: Remove?
type LoginArgs struct {
	Email    string
	Password string
}
type LoginCommand struct {
	Args LoginArgs
	fs   *flag.FlagSet
}

func (c *LoginCommand) Name() string {
	return "login"
}

func (c *LoginCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
	c.fs.StringVar(&c.Args.Email, "email", "", "Email address")
	c.fs.StringVar(&c.Args.Password, "password", "", "Password")
}

func (c *LoginCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *LoginCommand) Usage() {
	c.fs.Usage()
}

func (c *LoginCommand) Help() {
	fmt.Println("login -email <email> -password <password>")
}

func (c *LoginCommand) Execute(s *State) error {
	identity, err := s.GetIdentity(c.Args.Email, c.Args.Password)
	if err != nil {
		return fmt.Errorf("error logging in: %w", err)
	}

	log.Printf("Logged in as: %s\n", identity.Email)
	// TODO: Generate an identity key
	// TPDP: Store the key in a session (file: .identity)
	return nil
}
