package main

import (
	"bytes"
	"fmt"

	database "github.com/umesshk/database-normalizer/internal"
)

func Normalize(phone string) string {

	var new_phone bytes.Buffer

	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			new_phone.WriteRune(ch)
		}
	}

	return new_phone.String()
}

func main() {
	fmt.Println("Helo World")
	database.ConnectDatabase()
}
