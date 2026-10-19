package main

//Go is not an OOP language, however embedded structs provide  a kind of data only inheritance that can be useful at times
/** type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}. **/

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
	
}


// this means that user is embedded
// sender is a user that has a rateLimitfield