package main

func getMonthlyPrice(tier string) int {
	// ?
// writing the conditional statement using if and else if

	// if tier == "basic"{
	// 	return (100 * 100)
	// } else if tier == "premium"{
	// 	return  (150 * 100)
	// } else if tier == "enterprise"{
	// 	return (500 * 100)

	// } else{
	// 	return 0
	// }

	// writing a switch statement
	switch tier{
	case "basic":
		return (100 * 100)
	case "premium":
		return (150 * 100)
	case "enterprise":
		return (500 * 100)
	default:
		return 0
	}
}
