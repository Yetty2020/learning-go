package main

import "fmt"

//You have 5000 in your wallet, write a function that adds 1000 bonus.

func main(){
	balance := 5000
	newBalance := addBonus(balance)
	fmt.Println(newBalance)

	



}

func addBonus(balance int)int{
	balance = balance + 1000
	
	return balance

}

