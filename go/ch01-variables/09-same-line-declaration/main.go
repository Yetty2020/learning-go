package main

import "fmt"

func main() {
	// this is how to handle multiple variable declaration on a single line
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)
}
