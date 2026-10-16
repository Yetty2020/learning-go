package main

import ("fmt"
 "os")

func getGrade(score int) (string, bool) {

	if score < 0 || score > 100 {
		return "", false
	}
	if score >= 70 {
		return "A", true
	}
	if score >= 60 {
		return "B", true
	}
	if score >= 50 {
		return "C", true
	}
	if score >= 45 {
		return "D", true
	}
	if score >= 40 {
		return "E", true
	}
	return "F", true

}

func main() {

	length := len(os.Args)
	
	if length == 1 {
		fmt.Fprintln(os.Stderr, "usage: grade <score>")
		os.Exit(1)
	}
	

	fmt.Println(os.Args[1])

}
