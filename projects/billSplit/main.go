package main

import "errors"


func calcTip(bill, tipPercent float64)(float64, error){

	if bill <= 0 {
		return 0, errors.New("bill has to be greater than 0")
	}
	 if tipPercent < 0 {
		return  0, errors.New("tip cannot be negative")
	 }

	 tipAmount:= (tipPercent/100) * bill

	 


	return tipAmount, nil
}

func main(){

}