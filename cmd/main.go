package main

import (
	"log"

	"github.com/joho/godotenv"
	database "github.com/umesshk/database-normalizer/internal"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error Loading env variable")
	}

	db, err := database.ConnectDatabase()

	if err != nil {
		log.Fatal("Error Connecting to Database", err)
	}

	err = db.ResetDatabase()

	if err != nil {
		log.Fatal("Error Reseting Database")
	}

	if err = db.CreatePhoneTable(); err != nil {
		log.Fatal("Error Creating Table", err)
	}

	if err = db.SeedData(); err != nil {
		log.Fatal("Error Inserting Data ", err)
	}

	phones, err := db.GetAllPhone()

	if err != nil {
		log.Fatal("Error Getting Data", err)
	}

	if err := db.UpdateDb(phones); err != nil {
		log.Fatal("Error Updating Data", err)
	}

	defer db.Close()

}
