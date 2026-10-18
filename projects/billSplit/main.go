package main

import ("errors"
"fmt"
"os"
"strconv")


func calcTip(bill, tipPercent float64)(float64, error){

	if bill <= 0 {
		return 0, errors.New("bill has to be greater than 0")
	}
	 if tipPercent < 0 {
		return  0, errors.New("tip cannot be negative")
	 }

	 tipAmount:= (tipPercent/100) * bill
	 totalBill:= tipAmount + bill

	 


	return totalBill, nil
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

	//fmt.Println(os.Args)

	length := len(os.Args)

	if length != 4{
		fmt.Fprintf(os.Stderr, "split: billSplitter <bill> <people> <tip>")
		os.Exit(1)

	}


	bill, billError := strconv.ParseFloat(os.Args[1], 64)
	people, peopleError := strconv.Atoi(os.Args[2])
	tip, tipError := strconv.ParseFloat(os.Args[3], 64)


	if billError != nil{
		fmt.Fprintln(os.Stderr, "bill should be a number")
		os.Exit(1)
	}
	if peopleError != nil{
		fmt.Fprintln(os.Stderr, "person has to be a whole number")
		os.Exit(1)
	}
	if tipError != nil{
		fmt.Fprintln(os.Stderr, "tip cannot be a negative number")
		os.Exit(1)



	}

	totalBill, err := calcTip(bill, tip)
	if err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	splitBill, err := calcSplitBill(totalBill, people)
	if err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("You have a total bill of ₦%.2f, with a tip of %.1f %% and you are to split it amongst %d people, which gives ₦%.2f per person \n", bill, tip, people,splitBill)



	
	

}