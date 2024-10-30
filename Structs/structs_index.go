package main

import (
	"fmt"
	"time"
)
type user struct{
	firstName string
	surname string
	dob string
	createAt time.Time
}
func main() {
	var userFirstName = getInfo("Enter first name: ")
	var userSurname = getInfo("Enter surname: ")
	var userDob = getInfo("Enter date of birth (MM/DD/YYYY) ")
	
	var appUser user
	//instance of user
	appUser = user{
		firstName: userFirstName,
		surname: userSurname,
		dob: userDob,
		createAt: time.Now(),
	}
	//short notation but must be arranged well. values can be umitted or be empty -.> null
	// appUser = user{
	// 	userFirstName,
	// 	userSurname,
	// 	userDob,
	// 	 time.Now(),
	// }

	printDetail(appUser)
}
func printDetail(u user)  {
	fmt.Println(u.firstName,u.surname,u.dob )
}
func getInfo(detailQuestion string) string {
	fmt.Print(detailQuestion)
	var text string
	fmt.Scan(&text)
	return text
}