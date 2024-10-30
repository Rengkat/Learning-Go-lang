package main

import "fmt"

func main() {
	var amount float64
	var year int
	var rate float64
	fmt.Print("Enter invested amount: ")
	fmt.Scan(&amount)
	fmt.Print("Enter invested year: ")
	fmt.Scan(&year)
	fmt.Print("Enter invested rate: ")
	fmt.Scan(&rate)
	interest := calculate(amount,year, rate)
	fmt.Printf("Your interest is: %.1f", interest)

}

func calculate(amount float64, year int, rate float64 ) float64  {
	answer := (amount * rate * float64(year))/1000
	return answer
}