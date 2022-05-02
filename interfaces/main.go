package main

import (
	"fmt"
)

type animal interface {
	Says() string
	NumberOfLegs() int
}

type dog struct {
	name string
	breed string
}

type gorila struct {
	name string
	color string
	numberOfTeeth int
}

func main() {
dog := dog{
	name: "Samson",
	breed: "German Shephered",
}
printInfo(&dog)

gorila := gorila{
	name: "Jojo",
	color: "grey",
	numberOfTeeth: 32,
}
printInfo(&gorila)
}

func printInfo(a animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func(d *dog) Says() string {
	return "woof"
}

func(d *dog) NumberOfLegs() int {
	return 4
}

func(g *gorila) Says() string {
	return "ugh"
}

func(g *gorila) NumberOfLegs() int {
	return 2
}