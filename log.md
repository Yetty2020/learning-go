## Daily learning diary

## Day 3 - 09/10/25
- Go - Ch02 - Lesson 1 to 

Learned
- Learnt how to constants and how they are declared in Go.
- Constants are declared with the const keyword and the walrus operator is not needed.
- Values declared as contants cannot be changed when initally declared.
- const pi = 3.14
- Leant that constants must be known at compile time. They cannot be computed at runtime like Javascript.
- For example, if you see the Time.now() function, the value can only be computed at runtime.
- This cannot be declared as a constant in Go.
- Go is faster and lightweight than interpreted languages like python, javascript, Ruby, PHP.
- However, in terms of execution speed, Go is much slower than other compiled languages like C and Rust.
- This is due to the Go runtime that is used for memory management


## Day 2 - 07/10/2026
- Go - Ch01 - Lesson 5 to 14.


Learned
- Go is a complied language that needs a compiler, so the computer can understand it.
- Go is a high level language that is converted into maching language for the computer.
- Two types of error in Go - Runtime error and compilation error.
- Compilation Error happens when the code is compiled. A code with compilation error won't build, so it wont get to production.
- Runtime errors happens when the program is running. 
- Compilation error can be due to invalid syntax.
- Two types of comment styles in Go: single line and multi line. Just the way it is written in TS.
- Go uses type sizes for each data type.
- Signed Integer - int (1, 8, 16,32 64)
- Unsigned Integer - unint (1, 8, 16,32 64)
- Signed Decimal number - float (32 64)
- Complex numbers - complex (64,128)
- What type should you use? Use the default type.
- Only use a specific type when you are conerned about performance and memory usage.
- Go is statically typed and this means that variable types are known before the code runs. Just how TS works.
- You cannot concatenate different data types
- You cannot delcare multuple variables on this same line , this is called multiple line declaration.
- This is the syntax : mileage, company := 80276, "Toyota".
- Small Memory Footprint
- Go programs are fairly lightweight. 
- Every Go program has a small amount of extra code that is included in the executable binary code called Go Runtime.
- The purpose is to clean up unused memory at runtime.
- learnt different string formatting in Go:
- print, printf, println and sprintf
- The rule of thumb is use Println for quick output and debugging, and Printf when you need control (decimals, exact layout). Use Sprintf when you need the string in a variable.






## Day 1 - 06/10/2026
- Go - Installed Go, Set up repo and ch01 varibles folder.
- Linux - 

Learned
- How to set up a Go program.
- I had to declare the Package main,then import "fmt", so i could print to the console and then declared the main function.
- fmt.println - prints to the console
- Learnt how to declare varables in Go, using the var keyword (this is the sad way)
- the syntax is: the var keyword, the variable name and then the variable type
- var greeting string 
- Learnt five different types of variable types
- string - sequence of characters
- int - signed integer
- bool - boolean value whether true or false
- float64 - decimal value
- byte 8 bits of data
- Learnt how to declare variables the GOATED way, using the walrus operator (:=)
- You do not have to declare the varible type with this. The syntax is greeting := "hello world"

Struggles with


Tomorrow
3n hours


