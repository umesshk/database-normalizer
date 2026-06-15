package main

import (
	"github.com/joho/godotenv"
	database "github.com/umesshk/database-normalizer/internal"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error Loading env variable")
	}

	database.ConnectDatabase()
}
