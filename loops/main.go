package main

import (
	"log"
)

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	animals := []string{"dog", "fish", "cat", "horse", "cow"}

	for _, animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "Pluto"
	animals2["fish"] = "Cookies"
	animals2["cow"] = "Prime biff"
	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"

	for index, line := range firstLine {
log.Println(index, ":", line )
	}

	type user struct {
		firstName string
		lastName string
		email string
		age int
	}
	var users []user
	users = append(users, user{"John", "Smith", "john_smith@mail.com", 27})
	users = append(users, user{"Lary", "Smith", "lary_smith@mail.com", 40})
	users = append(users, user{"Dio", "Smith", "dio_smith@mail.com", 19})
	users = append(users, user{"Patrick", "Smith", "patrick_smith@mail.com", 25})

	for _, user := range users {
		log.Println(user.firstName, user.lastName, user.email, user.age)
	}
}