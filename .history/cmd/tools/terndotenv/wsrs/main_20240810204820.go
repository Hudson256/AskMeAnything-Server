package main

import "fmt"

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgx.New(ctx, fmt.Sprintf(""))
}
