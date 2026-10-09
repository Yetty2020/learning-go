package main


 import "fmt"

 func main(){
	//contants in Go must be known at compile time, and they cannot be changed at runtime.
	//They are usually declared with a static value
	//Unlike javascript, where you can declare a constant that can only be computed at run-time
	// for example, the current time can only be known when the program is running, so it cannot be declared as a contant.
	// to find the current time, this is done using time.Now() 

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = minutesInHour * secondsInMinute

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
 }