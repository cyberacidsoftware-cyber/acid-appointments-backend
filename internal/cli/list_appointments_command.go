package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

type ListAppointmentsArgs struct {
	Date string
}

type ListAppointmentsCommand struct {
	Args ListAppointmentsArgs
	fs   *flag.FlagSet
}

func (c *ListAppointmentsCommand) Name() string {
	return "list-appointments"
}

func (c *ListAppointmentsCommand) Init() {
	c.fs = flag.NewFlagSet(c.Name(), flag.ExitOnError)
	c.fs.StringVar(&c.Args.Date, "date", "", "Date to list appointments for")
}

func (c *ListAppointmentsCommand) Parse(args []string) {
	c.fs.Parse(args)
}

func (c *ListAppointmentsCommand) Usage() {
	c.fs.Usage()
}

func (c *ListAppointmentsCommand) Help() {
	fmt.Println("list-appointments -date <date>")
}

func (c *ListAppointmentsCommand) Execute(s *State) error {
	date := pgtype.Date{}
	if err := date.Scan(c.Args.Date); err != nil {
		return fmt.Errorf("invalid date: %w", err)
	}

	appointments, err := s.q.ListAppointmentsByDate(s.ctx, date)
	if err != nil {
		return fmt.Errorf("error listing appointments: %w", err)
	}

	for _, a := range appointments {
		fmt.Printf("Appointment %s\n", a.ID)
		date, err := a.Date.Value()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  Date: %s\n", date)
		time, err := a.Time.Value()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  Time: %s\n", time)
	}

	if len(appointments) == 0 {
		fmt.Println("No appointments found for date:", c.Args.Date)
	} else {
		fmt.Printf("Found %d appointments for date: %s\n", len(appointments), c.Args.Date)
	}

	return nil
}
