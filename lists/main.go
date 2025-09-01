package main

import "fmt"

func main() {
	names := make([]string, 2, 5)

	//append still adds two more to the slice making the size 4
	//but if we add cap 5 its not allocating new array
	names = append(names, "test")
	names = append(names, "test2")
	fmt.Println(names)

	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.8
	courseRatings["react"] = 5.0
	courseRatings["angular"] = 5.0
	fmt.Println(courseRatings)
	//no memory allocation
	//if add 4th item to the map it will need to make allocation
	//courseRatings["allocation"] = 1.0

	for index, value := range names {
		fmt.Println("Index", index)
		fmt.Println("Value", value)
		//index 0 and 1 will be empty
	}

	for key, value := range courseRatings {
		fmt.Println("Key", key)
		fmt.Println("Value", value)
	}
}
