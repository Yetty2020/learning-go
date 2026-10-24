package main

func bulkSend(numMessages int) float64 {

	total := 0.00

	for i := 0; i < numMessages; i++{
		total += 1.0 + float64(i) * 0.01
	}

	return total
}
