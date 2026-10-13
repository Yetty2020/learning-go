package main

func adder() func(int) int {

	var sum int

	return func(newSum int) int{
		sum += newSum
		return sum
	}
	
}
