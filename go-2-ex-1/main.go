package main

import "fmt"

type FullName struct {
	firstName string
	lastName string
}

// TODO: declare a structure for birth date

type Birthdate struct {
	dayOfBirth int
	monthOfBirth int
	yearOfBirth int
}

type Profile struct {
	// TODO: embed full name and birth date information

	Name FullName
	birthdate Birthdate
	NumberOfSiblings byte
	ZodiacSign       rune
	heightInMeters	float32
}

func main() {
	var me = Profile{
		Name: FullName{
			firstName: "Patrick",
			lastName: "Meier",
		},

		birthdate: Birthdate {
			dayOfBirth: 20,
			monthOfBirth: 10,
			yearOfBirth: 2007,
		},
		// TODO: set name and birth date information
		NumberOfSiblings: 3,
		ZodiacSign:	'\u264E',
		heightInMeters: 1.87,
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings = me.NumberOfSiblings + 1
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
