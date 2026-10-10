package main

// import "fmt"

func main(){

	// Normal way: length lives in the whole function
// length := getLength(email)
// if length < 10 {
// 	fmt.Println("too short")
// }
// fmt.Println(length) // ✅ still works here, even though you don't need it

//Use this because it makes the code shorter and cleaner.

// Initial statement: length only lives inside the if/else
// if length := getLength(email); length < 10 {
// 	fmt.Println("too short")
// } else {
// 	fmt.Println(length) // ✅ works in else too
// }
// fmt.Println(length) // error: undefined: length
}