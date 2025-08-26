package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	//fmt.Print("Investment amout: ")
	outputText("Investment amout: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Exptected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
