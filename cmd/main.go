package main

import (
	"fmt"
	"github.com/joho/godotenv"
	database "github.com/umesshk/database-normalizer/internal"
	"regexp"
)

func Normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}

// func Normalize(phone string) string {
//
// 	var new_phone bytes.Buffer
//
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			new_phone.WriteRune(ch)
// 		}
// 	}
//
// 	return new_phone.String()
// }

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error Loading env variable")
	}

	fmt.Println("Helo World")
	database.ConnectDatabase()
}
