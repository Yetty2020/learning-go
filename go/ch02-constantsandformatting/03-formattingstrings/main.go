package main

import "fmt"

func main(){
	const name = "Fatiha"
	const age = 29
	// learning how to format strings in go
	// fmt.print, this prints the string to the console without adding a new line
	fmt.Print("Hello, World!\n")
	//fmt.println - this prints the values as they are and adds a new line at the end of the string, thats what the ln does
	fmt.Println("HI, my name is",name, "and i am", age)
	//fmt.printf - this prints a formatted string where you control everthing, it also takes in placeholders
	fmt.Printf("Hi, my name is %s and I am %d years old", name, age)

	//Default representation for each placeholder:
	// %v - default representation of the value
	// %T - type of the value
	// %d - decimal representation of an integer
	// %f - decimal representation of a floating-point number
	// %s - string representation of a string
	// %t - boolean representation of a boolean value

	fmt.Printf("\nThe default representation of name is: %v", name)
	fmt.Printf("\nThe type of name is: %T", name)
	fmt.Printf("\nThe decimal representation of age is: %d", age)
	fmt.Printf("\nThe string representation of name is: %s", name)
	fmt.Printf("\nThe boolean representation of true is: %t", true)

	//fmt.Sprintf - this returns a formatted string as a variables instead of printint it to the console.
	info := fmt.Sprintf("\nHi, my name is %s and I am %d years old", name, age)
	fmt.Println(info)
}