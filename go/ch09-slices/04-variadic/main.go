package main

//Variadic means “takes any number of arguments.” You already know this from TS, where it’s called rest parameters.
//func sum(nums ...int) int {
//nums is the name, and ...int means “any number of ints.” The caller can pass zero, one or many:
//The ... tells it. When Go sees ... in front of a type in the parameter list, the rule is:

//“Collect all the arguments of this type into a slice.”

//numbers := []int{1, 2, 3}
//sum(numbers)   // ❌ error!
// to fix this error, you have to pass it as individual parameters

//sum(numbers...) 
//The ... after the slice means “take everything out of this slice and pass it as separate arguments.
//sum(1, 2, 3)
func sum(nums ...int) int {

	total := 0

	for i := 0; i < len(nums); i++{
		total += nums[i]
		
	}
	return total
	
}
