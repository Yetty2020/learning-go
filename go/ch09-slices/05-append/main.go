package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	newArray := []float64{}

	for i := 0; i < len(costs); i++{
		if costs[i].day == day{
			 newArray = append(newArray, costs[i].value)
			
		}
	}
	return newArray
	
	
}
