package main

func getMonthlyPrice(tier string) int {
	// ?


	if tier == "basic"{
		return (100 * 100)
	} else if tier == "premium"{
		return  (150 * 100)
	} else if tier == "enterprise"{
		return (500 * 100)

	} else{
		return 0
	}
}
