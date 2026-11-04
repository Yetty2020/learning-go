package main

// to create a slice directly instead of trying to get it from an array, you can use the make keyword
// mySlice := make([]int,3,4)
// func make([]T, len, cap) []T
// capacity parameter is usually left out because ti equals the length

func getMessageCosts(messages []string) []float64 {

	length := len(messages)
	
	costs := make([]float64, length)

	for i := 0; i< length; i++{
		costs[i] = float64(len(messages[i])) * 0.01
		
	}

	return costs
	// ?
}
