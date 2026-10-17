package main


func getStatus(percentage float64)string{

	if percentage >= 100{
		return "Exhausted"
	} 
	if percentage >= 80{
		return "Running low"
	} 
	return "OK"

	
}

func main(){

}