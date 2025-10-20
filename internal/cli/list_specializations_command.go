package cli

import (
	"flag"
	"fmt"

	"backend/internal/usecase"
)

type ListSpecializationsArgs struct {
}

type ListSpecializationsCommand struct {
	fs *flag.FlagSet
}

func (c *ListSpecializationsCommand) Name() string {
	return "list-specializations"
}

func (c *ListSpecializationsCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
}

func (c *ListSpecializationsCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *ListSpecializationsCommand) Usage() {
	c.fs.Usage()
}

func (c *ListSpecializationsCommand) Help() {
	fmt.Println("list-specializations")
}

func (c *ListSpecializationsCommand) Execute(s *State) error {
	specializations, err := usecase.ListSpecializations(s)
	if err != nil {
		return err.Error
	}

	if len(specializations) == 0 {
		fmt.Println("No specializations found")
	}

	for _, s := range specializations {
		fmt.Printf("%s | %s\n", s.ID.String(), s.Name)
	}

	return nil
}
