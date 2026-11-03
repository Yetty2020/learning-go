package main
//an array is a group of values of the same type stored under the same name, with a fixed size.
//[4]string
//this says what goes into the slot, which is strings
//4 is the no of slots
//var myInts [10]int
// this makes 10 slots for null values of integer type
// to delcare the array with values, you use curly braces
// primes := [4]int{2,2,3,4}
// to get values by index 
// primes[3] = 4

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {

		message := [3]string{primary, secondary, tertiary}

	var costs [3]int // tthis creates an empty array of integers
	total := 0 // this is the running accumulator
	for i:= 0; i<len(message); i++{

		total+= len(message[i])
		costs[i] = total

	}



	
	
	return message, costs
}
