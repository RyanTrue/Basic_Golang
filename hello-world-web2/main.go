package main

import "log"

func main() {
isTrue := true
	if isTrue == true {
		log.Println("Is true is", isTrue)
	} else {
		log.Println("Is true is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("It's", cat)
	} else {
		log.Println("It's not", cat)
	}

	myNum := 12
	isTrue = false

	if myNum == 12 && !isTrue {
		log.Println("My num", myNum)
	} else {
		log.Println("My num not", myNum)
	}

	myVar := "dog"

	switch myVar {
	case "dog":
		log.Println("dog is set to dog")
	case "cat":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default: 
	log.Println("cat is something else")
	}
}