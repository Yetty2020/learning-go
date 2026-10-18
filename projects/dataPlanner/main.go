package main

import "errors"

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

}
