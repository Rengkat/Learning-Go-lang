package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var year float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate / 100,year)
	fmt.Println(futureValue)

}