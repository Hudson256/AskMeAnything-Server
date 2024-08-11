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

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("WSRS_DATABASE_USER"), os.Getenv("WSRS_DATABASE_PASSWORD"), os.Getenv("WSRS_DATABASE_HOST"), os.Getenv("WSRS_DATABASE_PORT"), os.Getenv("WSRS_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}

	defer pool.close()

	if err := pool.Ping(ctx); err != nil {

	}
}
