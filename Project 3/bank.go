package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
const accountBalance = "balance.txt"
func readBalanceFile() (float64, error)  {
	data, err := os.ReadFile(accountBalance) //ths must be converted to string first
	if err != nil{
		return 100, errors.New("failed to find balance file")
	}
	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)
	if err != nil{
		return 1000, errors.New("")
	}
	return balance, nil
}
func writeBalanceToFile(balance float64)  {
	balanceTxt :=fmt.Sprint(balance)
	os.WriteFile(accountBalance,[]byte(balanceTxt),0644)
}
func main() {
	accountBalance, err := readBalanceFile()
	if err != nil{
		fmt.Println("ERROR")
		fmt.Print(err)
		fmt.Println("----------------")
		//you can exit the program once there is error
		//1. by using return
		//3. by using panic function
		panic("Sorry, can't continue")
	}
	fmt.Println("Welcome to Go bank")
	for {
		
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check account balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance )
		case 2:
			var deposited float64
			fmt.Print("Enter deposited amount: ")
			fmt.Scan(&deposited)
			accountBalance +=float64(deposited)
			fmt.Println("Cash deposited. Your new account is:", accountBalance)
			writeBalanceToFile(float64(accountBalance))
		case 3:
			var withdrawAmount float64
			fmt.Print("Amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0{
				fmt.Println("Invalid amount. Amount must be greater than 0")
				continue
			}
			if withdrawAmount > float64(accountBalance){
				fmt.Println("Insufficient balance")
				continue
			}
			accountBalance -= float64(withdrawAmount)
			fmt.Println("Cash withdrawal successful. Your new balance is:", accountBalance)
			writeBalanceToFile(float64(accountBalance))
		default:
		fmt.Println("Goodbye...")
		return
		// break
		}


		// if choice == 1{
		// 	fmt.Println("Your account balance is:",accountBalance )
		// 	} else if choice == 2{
		// 		var deposited float64
		// 		fmt.Print("Enter deposited amount: ")
		// 		fmt.Scan(&deposited)
		// 		accountBalance +=int(deposited)
		// 		fmt.Println("Cash deposited. Your new account is:", accountBalance)	
		// } else if choice == 3{
		// 	var withdrawAmount float64
		// 	fmt.Print("Amount to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0{
		// 		fmt.Println("Invalid amount. Amount must be greater than 0")
		// 		continue
		// 	}
		// 	if withdrawAmount > float64(accountBalance){
		// 		fmt.Println("Insufficient balance")
		// 		return
		// 	}
		// 	accountBalance -= int(withdrawAmount)
		// 	fmt.Println("Cash withdrawal successful. Your new balance is:", accountBalance)
		// }else{
		// 	fmt.Println("Goodbye...")
		// 	break
		// }
	}
}