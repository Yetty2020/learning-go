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

func calcSplitBill(total float64, people int)(float64, error){

	if total <= 0{
		return 0, errors.New("total must be greater than 0")
	}
	if people <=0{
		return 0, errors.New("need at least 1 person to split the bill")
	}

	eachShare := total / float64(people)
	



	return eachShare, nil
}

func main(){

}