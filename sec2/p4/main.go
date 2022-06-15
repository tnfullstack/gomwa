package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Thanh",
		LastName:    "Nguyen",
		PhoneNumber: "555-666-7789",
		Age:         50,
	}

	log.Println(user.FirstName, user.LastName)

}
