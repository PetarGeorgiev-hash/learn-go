package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	balance, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000.0, errors.New("failed to find balance")
	}
	balanceText := string(balance)
	floatBalance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000.0, errors.New("failed to parse balance")
	}
	return floatBalance, nil
}

func writeToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0664)
}

func main() {
	fmt.Println("Welcome to GO Bank")
	accBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Can't continue sorry")
	}
	for {
		fmt.Println("What do you wish to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println(accBalance)
		case 2:
			fmt.Print("Deposit amonut: ")
			var depositAmonut float64
			fmt.Scan(&depositAmonut)
			if depositAmonut <= 0 {
				fmt.Println("Invalid amonut")
				continue
			}
			accBalance += depositAmonut
			fmt.Println("Account balance updated: ", accBalance)
			writeToFile(accBalance)
		case 3:
			fmt.Print("Withdraw amaunt: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accBalance {
				fmt.Println("Can't withdraw more than your account balance")
				continue
			}

			accBalance -= withdrawAmount
			fmt.Println("Account balance updated: ", accBalance)
			writeToFile(accBalance)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
