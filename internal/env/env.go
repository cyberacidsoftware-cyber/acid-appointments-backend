package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Optional
const (
	APP_ENVIRONMENT = "APPOINTMENTS_ENVIRONMENT"
)

// Required
const (
	DATABASE_USER     = "APPOINTMENTS_DATABASE_USER"
	DATABASE_PASSWORD = "APPOINTMENTS_DATABASE_PASSWORD"
	DATABASE_HOST     = "APPOINTMENTS_DATABASE_HOST"
	DATABASE_PORT     = "APPOINTMENTS_DATABASE_PORT"
	DATABASE_NAME     = "APPOINTMENTS_DATABASE_NAME"
	JWT_SECRET        = "APPOINTMENTS_JWT_SECRET"
)

func check() error {
	envNames := []string{
		DATABASE_USER,
		DATABASE_PASSWORD,
		DATABASE_HOST,
		DATABASE_PORT,
		DATABASE_NAME,
		JWT_SECRET,
	}
	for _, name := range envNames {
		if Get(name) == "" {
			return fmt.Errorf("missing environment variable %s", name)
		}
	}
	return nil
}

func Load() error {
	if _, err := os.Stat(".env"); err == nil {
		return godotenv.Load()
	}
	return check()
}

func Get(name string) string {
	return os.Getenv(name)
}

func GetDatabaseConnection() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv(DATABASE_USER),
		os.Getenv(DATABASE_PASSWORD),
		os.Getenv(DATABASE_HOST),
		os.Getenv(DATABASE_PORT),
		os.Getenv(DATABASE_NAME),
	)
}
