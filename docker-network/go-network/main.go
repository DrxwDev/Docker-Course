package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load the ENV file: %v\n", err)
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, os.Getenv("DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create pool: %v\n", err)
		return
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Database Not reachable: %v", err)
		return
	}

	fmt.Println("Database UP and Running....")
}
