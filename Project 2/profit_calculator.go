package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	printText("Revenue: ")
	fmt.Scan(&revenue)

	printText("Expenses: ")
	fmt.Scan(&expenses)

	printText("Tax rate: ")
	fmt.Scan(&taxRate)

	ebt :=revenue - expenses
	profit := ebt * (1- taxRate/100)
	ratio := ebt/profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	fmt.Printf(`the ration is: %.2f`,ratio) //2d.p-> use Printf with %
	// Text(4,6)
}
func printText(txt string)  {
	fmt.Print(txt)
}
func calculateProfit(revenue float64, expenses float64, taxRate float64)  {
	ebt :=revenue - expenses
	profit := ebt * (1- taxRate/100)
	return profit
	
}