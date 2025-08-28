package main

import (
	"fmt"
)

func main() {
	age := 32
	agePointer := &age
	fmt.Println(*agePointer)
	fmt.Println(editAgeToAdult(agePointer))
}

func editAgeToAdult(age *int) int {
	return *age - 18
}
