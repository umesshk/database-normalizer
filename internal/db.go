package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/umesshk/database-normalizer/normalizer"
)

type Phone struct {
	Id     int
	Number string
}

func ConnectDatabase() {

	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_DATABSE")
	)

	conn_url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn_url)

	if err != nil {
		fmt.Println("Error connecting")
		panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error Reaching database")
		panic(err)
	}

	fmt.Println("Connected Succefully")

	if err := CreatePhoneTable(db); err != nil {
		fmt.Println("Error Creating Table ")
		panic(err)
	}
	//
	// phone_numbers := []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}
	//
	// for _, ph := range phone_numbers {
	// 	i, err := InsertData(db, ph)
	// 	fmt.Println("id=", i)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	//

	phones, err := GetAllPhone(db)

	if err != nil {
		log.Fatal(err)
	}

	for i := range phones {
		npH := normalizer.Normalize(phones[i].Number)
		phones[i].Number = npH
	}

	err = UpdateDb(db, phones)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}

func CreatePhoneTable(db *sql.DB) error {

	query := `CREATE TABLE IF NOT EXISTS phone_numbers ( 
					id   SERIAL PRIMARY KEY,
					value varchar(225)
	);`

	_, err := db.Exec(query)

	return err

}

func InsertData(db *sql.DB, phone_num string) (int, error) {
	query := `INSERT INTO phone_numbers (value) VALUES($1) RETURNING id `

	var id int
	err := db.QueryRow(query, phone_num).Scan(&id)

	return id, err
}

func GetPhone(db *sql.DB, id int) (error, string) {

	var phone_number string

	query := `SELECT value from phone_numbers WHERE id=$1`

	err := db.QueryRow(query, id).Scan(&phone_number)

	return err, phone_number
}

func GetAllPhone(db *sql.DB) ([]Phone, error) {

	row, err := db.Query(`SELECT id,value from phone_numbers `)
	if err != nil {
		return nil, err
	}

	var p []Phone

	for row.Next() {
		var ph Phone
		err := row.Scan(&ph.Id, &ph.Number)

		if err != nil {
			return nil, err
		}

		p = append(p, ph)

	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}

	return p, nil
}

func UpdateDb(db *sql.DB, phones []Phone) error {

	for _, ph := range phones {

		_, err := db.Exec(`UPDATE  phone_numbers set value= $1 where id =$2`, ph.Number, ph.Id)

		if err != nil {
			return err
		}

	}

	fmt.Println("Updated Tables Succefully...")

	return nil

}
