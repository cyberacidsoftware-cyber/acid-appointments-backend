package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"backend/internal/env"

	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	connectionString := env.GetDatabaseConnection()
	var pool *pgxpool.Pool
	var err error

	maxRetries := 5
	waitTime := 2 * time.Second
	maxTime := 10 * time.Second

	for i := 0; i < maxRetries; i++ {
		pool, err = pgxpool.New(ctx, connectionString)
		if err != nil {
			log.Printf("Failed to create DB pool (attempt %d/%d): %v", i+1, maxRetries, err)
		} else if err = pool.Ping(ctx); err != nil {
			log.Printf("DB not ready yet (attempt %d/%d): %v", i+1, maxRetries, err)
		} else {
			// success
			break
		}

		time.Sleep(waitTime)
		if waitTime > maxTime {
			waitTime = maxTime
		} else {
			waitTime *= 2 // exponential backoff
		}
	}

	if err != nil {
		log.Fatal("Could not connect to database after retries:", err)
	}
	defer pool.Close()
	var auth *jwtauth.JWTAuth
	// Define a secret key for signing the JWT tokens
	jwtSecret := env.Get(env.JWT_SECRET)
	auth = jwtauth.New("HS256", []byte(jwtSecret), nil)

	handler := NewApi(pool, auth)

	go func() {
		log.Println("Server started at :8080")
		if err := http.ListenAndServe(":8080", handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
