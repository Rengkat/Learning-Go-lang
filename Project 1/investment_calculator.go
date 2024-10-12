package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var year = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate / 100,float64(year))
	fmt.Println(futureValue)

}