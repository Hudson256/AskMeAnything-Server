package main

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
