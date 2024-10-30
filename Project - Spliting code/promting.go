package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func prompt() {
	//external pakage
	fmt.Println(randomdata.Female)
	var amount float64
	var year int
	var rate float64
	fmt.Print("Enter invested amount: ")
	fmt.Scan(&amount)
	fmt.Print("Enter invested year: ")
	fmt.Scan(&year)
	fmt.Print("Enter invested rate: ")
	fmt.Scan(&rate)
}
