package main

import (
	"fmt"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgx.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s,
		os.Getenv("WSRS_DATABASE_USER")
		))
}
