package main

import (
	"log"
	"time"
)

type user struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := user {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "8 800-55-35-35",
		Age: 35,
	}
	log.Println(user.FirstName, user.LastName, user.BirthDate, user.PhoneNumber, user.Age)
}