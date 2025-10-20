package main

import (
	"log"
	"os"
	"os/exec"

	"backend/internal/env"
)

func main() {
	if err := env.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: w_env <command> [args...]")
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	cmd.Run()
}
