package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createTransformerFunc(2)
	triple := createTransformerFunc(3)

	transformed := transformNumbers(&numbers, double)
	transformed = transformNumbers(&numbers, triple)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformerFunc(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
