package cli

import (
	"backend/internal/usecase"
	"flag"
	"fmt"
	"log"
)

type ResetDbCommand struct {
	fs *flag.FlagSet
}

func (c *ResetDbCommand) Name() string {
	return "reset-db"
}

func (c *ResetDbCommand) Init() {
}

func (c *ResetDbCommand) Parse(args []string) {
}

func (c *ResetDbCommand) Usage() {
	c.fs.Usage()
}

func (c *ResetDbCommand) Help() {
	fmt.Println("reset-db")
}

func (c *ResetDbCommand) Execute(s *State) error {
	err := usecase.ResetDb(s)
	if err != nil {
		return err.Error
	}
	log.Printf("Database reset\n")
	return nil
}
