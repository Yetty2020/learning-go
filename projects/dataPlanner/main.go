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
func calcUsage(plan, usage float64)(float64, float64, error){
	return 0, 0, nil
}

func main(){

}