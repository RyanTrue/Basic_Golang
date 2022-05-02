package main

import "log"

type myStruct struct {
	firstName string
}

func (m *myStruct) printFirstName() string {
return m.firstName
}

func main() {
var myVar myStruct
myVar.firstName = "John"

myVar2 := myStruct{
	firstName: "Mary",
}

log.Println("myVar is set tp", myVar.printFirstName())
log.Println("myVar2 is set tp", myVar2.printFirstName())
}