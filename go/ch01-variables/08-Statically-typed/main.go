package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"

	// don't edit below this line
	// this is to conconatnate two strings, you cannnot concatenate a string and an int
	
	fmt.Println("Authorization: Basic", username+":"+password)
}
