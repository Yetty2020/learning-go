package main


//learning about structures in Go, they are used like dictionaries in python or objects in TS.
// It is used to represent structured data
// For example, to represent a car
/** type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
}.  **/


// And then to create a car

/** myCar := car{
	brand:   "Toyota",
	model:   "Camry",
	doors:   4,
	mileage: 5000,
}
	**/



type messageToSend struct {
	// ?

	phoneNumber int
	message string
}
