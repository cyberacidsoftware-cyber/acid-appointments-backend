package main

import (
	"backend/internal/cli"
	"os"
)

func main() {
	// var programName = os.Args[0]
	programName := "go run cmd/cli/main.go"
	args := os.Args[1:]
	cli.NewCli(programName).Run(args)
}
