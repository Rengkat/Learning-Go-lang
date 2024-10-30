package main

import (
	"errors"
	"fmt"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue,err :=printText("Revenue: ")
	if err != nil{
		fmt.Print(err)
	}
	expenses,err:=printText("Expenses: ")
	taxRate,err:=printText("Tax rate: ")
	
	ebt, profit, ratio := calculateProfit(revenue,expenses,taxRate)
	// ebt :=revenue - expenses
	// profit := ebt * (1- taxRate/100)
	// ratio := ebt/profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	fmt.Printf(`the ration is: %.2f`,ratio) //2d.p-> use Printf with %
	// Text(4,6)
}
func printText(txt string) (float64, error){
	var userInput float64
	fmt.Print(txt)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Number must be positive value")
	}
	return userInput, nil
}

func calculateProfit(revenue float64, expenses float64, taxRate float64) (float64,float64,float64){
	ebt :=revenue - expenses
	profit := ebt * (1- taxRate/100)
	ratio := ebt/profit
return ebt, profit,ratio
	
}
//alternatively
// func calculateProfitAlt(revenue float64, expenses float64, taxRate float64) (ebt float64,profit float64,ratio float64){
// 	ebt =revenue - expenses
// 	profit = ebt * (1- taxRate/100)
// 	ratio = ebt/profit
// 	return
	
// }