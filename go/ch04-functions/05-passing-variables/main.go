package main

import "fmt"

func main(){
	x := 5
	x = increment((x))
	
	fmt.Println(x)
	

}

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int

	lastMonthBill = getBillForMonth( costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth( costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth( costPerSend, messagesSent int) int {
	bill := costPerSend * messagesSent
	return bill
	
}

func increment(x int)int{
	x++
	return x
}