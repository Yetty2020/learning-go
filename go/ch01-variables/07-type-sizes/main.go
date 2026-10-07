package main

import "fmt"

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)
	//this converts the the float number to a signed integer

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
