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

type Database struct {
	DB *sql.DB
}

func ConnectDatabase() (*Database, error) {

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
		return &Database{}, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error Reaching database")
		panic(err)
	}

	fmt.Println("Connected Succefully...")

	return &Database{db}, nil

}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) SeedData() error {

	fmt.Println("Inserting Data...")

	phone_numbers := []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}

	for _, ph := range phone_numbers {
		i, err := d.InsertData(ph)
		fmt.Println("Inserted Data with id=", i)
		if err != nil {
			return err
		}
	}

	fmt.Println("Data Inserted...")

	return nil

}

func (d *Database) CreatePhoneTable() error {
	query := `CREATE TABLE IF NOT EXISTS phone_numbers ( 
					id   SERIAL PRIMARY KEY,
					value varchar(225)
	);`

	_, err := d.DB.Exec(query)

	if err == nil {
		fmt.Println("Created Table Phone phone_numbers...")
	}
	return err

}

func (d *Database) InsertData(phone_num string) (int, error) {

	query := `INSERT INTO phone_numbers (value) VALUES($1) RETURNING id `

	var id int
	err := d.DB.QueryRow(query, phone_num).Scan(&id)

	return id, err
}

func (d *Database) GetPhone(id int) (error, string) {

	var phone_number string

	query := `SELECT value from phone_numbers WHERE id=$1`

	err := d.DB.QueryRow(query, id).Scan(&phone_number)

	return err, phone_number
}

func (d *Database) GetAllPhone() ([]Phone, error) {

	fmt.Println("Fetching Alll Numbers form Database...")
	row, err := d.DB.Query(`SELECT id,value from phone_numbers `)
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

	fmt.Println("Data Fetched...")
	return p, nil
}

func (d *Database) CheckPhone(ph_num string) (*Phone, error) {
	fmt.Printf("Checking Phone Number : %s...", ph_num)
	var phone_number Phone

	query := `SELECT id , value from phone_numbers WHERE value=$1`

	err := d.DB.QueryRow(query, ph_num).Scan(&phone_number.Id, &phone_number.Number)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &phone_number, err
}

func (d *Database) UpdateDb(phones []Phone) error {

	fmt.Println("Updating Database...")

	for _, ph := range phones {

		new_ph := normalizer.Normalize(ph.Number)

		if new_ph == ph.Number {
			fmt.Println("No operation performed...", new_ph)
		} else {

			num, err := d.CheckPhone(new_ph)

			if err != nil {
				panic(err)
			}

			if num != nil {

				if num.Id != ph.Id {

					err := d.DeleteRecord(ph.Id)
					if err != nil {
						panic(err)
					}

				}
			}

			_, err = d.DB.Exec(`UPDATE  phone_numbers set value= $1 where id =$2`, new_ph, ph.Id)

			if err != nil {
				return err
			}
			fmt.Println("Updated Tables Succefully...", new_ph)

		}

	}

	return nil

}

func (d *Database) DeleteRecord(id int) error {

	query := `DELETE FROM phone_numbers WHERE id=$1`

	_, err := d.DB.Exec(query, id)

	if err != nil {
		return err
	}

	fmt.Println("Deleted Phone Number with id ", id)

	return nil

}

func (d *Database) ResetDatabase() error {
	query := `DROP TABLE phone_numbers`

	_, err := d.DB.Exec(query)

	fmt.Println("Database Reseted...")

	return err

}
