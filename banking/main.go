package main

import (
	"fmt"

	"github.com/PetarGeorgiev-hash/learning-go/file"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to GO Bank")
	accBalance, err := file.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		fmt.Println("Can't continue sorry")
	}
	presentOptions()

	for {

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
			file.WriteFloatToFile(accBalance, accountBalanceFile)
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
			file.WriteFloatToFile(accBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
