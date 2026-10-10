package main


import "fmt"

func main(){

	//When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.
fmt.Print("Hello")
}

func addToDatabase(hp, damage int) {
  // ...
}

func addToDatabases(hp, damage int, name string, level int) {
  // ?
}