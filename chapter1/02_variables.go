package main

import "fmt"

func main() {
	//using the walrus operator to declare and assign a varible
	messageStart := "You are learning GO!"
	fmt.Println(messageStart)

	//setting different variable types
	var name string
	name = "Fatihah"
	var isTrue bool
	isTrue = true
	var floatNum float64
	floatNum = 3.14
	var newAge int
	newAge = 30
	fmt.Println("I just clocked", newAge)
	fmt.Println("I am", name, "and I am", floatNum, "years old.", "This information is", isTrue)

}
