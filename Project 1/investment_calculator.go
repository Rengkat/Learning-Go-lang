package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate= 2.0
	var investmentAmount float64 = 1000
	 expectedReturnRate := 5.5
	var year float64 = 10
fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("year: ")
	fmt.Scan(&year)

	fmt.Print("expected rate return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate / 100,year)
	futureRealValue := futureValue/math.Pow(1+inflationRate/100,year)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	age, broAge := 25, 14
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	
	sumOfAge := age + broAge
	fmt.Println(sumOfAge)

}