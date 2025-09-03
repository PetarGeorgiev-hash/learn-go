package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubleNumber := transformNumbers(&numbers, double)
	fmt.Println(doubleNumber)
	trippleNumbers := transformNumbers(&numbers, tripple)
	fmt.Println(trippleNumbers)

	moreNumbers := []int{5, 6, 7, 8}

	transofrmFn1 := getTransformerFunction(&numbers)
	transofrmFn2 := getTransformerFunction(&moreNumbers)

	transformedNumber := transformNumbers(&moreNumbers, transofrmFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transofrmFn2)

	fmt.Println(transformedNumber)
	fmt.Println(moreTransformedNumbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func tripple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return tripple
	}
}
