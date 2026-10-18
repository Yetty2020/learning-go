package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getStatus(percentage float64) string {

	if percentage >= 100 {
		return "Exhausted"
	}
	if percentage >= 80 {
		return "Running low"
	}
	return "OK"

}
func calcUsage(plan, usage float64) (float64, float64, error) {

	if plan <= 0 {

		return 0, 0, errors.New("You do not have an active data plan")

	}
	if usage < 0 {
		return 0, 0, errors.New("Usage Plan cannot be negative")
	}

	remainder := plan - usage
	percUsed := (usage / plan) * 100

	if remainder < 0 {
		remainder = 0
	}
	return remainder, percUsed, nil

}

func main() {
	// this is used to seet the output of what we type
	// the first index is the file adress, followed by what we type
	// fmt.Println(os.Args)
	length := len(os.Args)

	// first error to handle when 2 inputs are not received
	if length != 3 {
		fmt.Fprintln(os.Stderr, "usage: dataPlanner <plan_gb> <used_gb>")
		os.Exit(1)
	}

	// error to handle conversion from strings to floating points
	//  plan, usage, error := strconv.ParseFloat(os.Args[])
	//  fmt.Println(plan, usage, error)

	plan, planErr := strconv.ParseFloat(os.Args[1], 64)
	usage, usageErr := strconv.ParseFloat(os.Args[2], 64)

	//fmt.Println(plan, planErr, usage, usageErr)

	if planErr != nil {
		fmt.Fprintln(os.Stderr, "Plan must be a number")
		os.Exit(1)
	}
	if usageErr != nil {
		fmt.Fprintln(os.Stderr, "Usage must be a number")
		os.Exit(1)

	}

	remainder, percentage, err := calcUsage(plan, usage)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}
	result := getStatus(percentage)
	fmt.Printf("You have used up %.1f %% of your data plan. You have %.1fGB left. Your current status is: %s\n", percentage, remainder, result)

}
