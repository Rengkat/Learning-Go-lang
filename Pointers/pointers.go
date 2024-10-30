package main

import "fmt"

func main() {
	age := 40 //regular variable
	agePointer := &age
	depointedAge := *agePointer
	fmt.Println(depointedAge)
	fmt.Println(getAdultYear(agePointer))

}

func getAdultYear(age *int) int {
	return *age - 18
}