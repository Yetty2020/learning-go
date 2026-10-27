## Daily learning diary
## Day 18 - 24/10/25


Error is a built in interface
You can write your own custom errors by specifying the error() method on the structs
Here's an example of a userError struct that implements the error interface:

type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

It can then be used as an error:

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}


- There is an errors package called errors.New(), you can just use this instead of overcomplicating matters and trying to write your own error interface


## Loops in Go
This is similar to how it is written in Ts, except that brackets are not used in Go.
This is the syntax
for INITIAL; CONDITION; AFTER{
  // do something
}

For example:

for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9

Rule of thumb:

You know the count up front: use a condition (lesson 1).
You find the end as you go: leave the condition empty and exit from inside (this lesson). You omit the conditon

## Day 17 - 23/10/25

## Type Assertions and Type Switches

**The problem:** when a function takes an interface, Go only knows the interface's methods. It forgets the real type, so type-specific fields aren't allowed.

```go
func getExpenseReport(e expense) (string, float64) {
	e.toAddress // ❌ not every expense has toAddress
}
```

Like a sealed parcel labelled "expense": I know it has a price (`cost()`), but I can't see what's inside.

### Type assertion: check for ONE type

```go
em, ok := e.(email) // "is e actually an email?"
if ok {
	return em.toAddress, em.cost() // ✅ safe, em is a real email
}
```

- `ok` is `true` if it is that type, and `false` if not
- `em` is the value as the real type. Only use it when `ok` is true.
- Same pattern as `n, err := strconv.Atoi(...)`: check first, then use.
- ⚠️ `em := e.(email)` without `ok` **panics** (crashes) if it's the wrong type.

### Type switch: check for SEVERAL types

```go
switch v := e.(type) {
case email:
	return v.toAddress, v.cost() // v is an email here
case sms:
	return v.toPhoneNumber, v.cost() // v is an sms here
default:
	return "", 0.0 // anything else
}
```

- `.(type)` is literal and only works inside a `switch`
- `v` becomes the right type in each case
- no `ok` and no `if` needed, because the switch already did the check
- no `break` needed, because cases don't fall through

### Which one to use?

| Checking... | Use |
|---|---|
| one type | assertion + `ok` |
| two or more types | type switch |

They do the same job, so pick **one**. Don't mix them.

### Mistakes I made

- Putting the fallback `return` too early, which made the code after it unreachable (like `exit 0` in the middle of a bash script)
- Using assertions *and* a switch together, which does the checking twice, and leaves `v` unused (a compile error)

### Links to what I know

- **TS:** like `if (x instanceof Email)` / `switch` with type narrowing. The `ok` version is a safe `instanceof`, not a blind `as`.
- **Linux:** like the `file` command, which checks what something *really* is instead of trusting its name. A type switch is like bash's `case ... in` with `*)` as the default.
- **`%T`** prints a value's type name, but only an assertion or switch lets you *use* the type.

## Day 16- 22/10/25
- Go Ch05 Interfaces in Go

This is the entire syntax

1. For the Interface

type NAME interface {
	method() returnType
}

2. struct
type TYPE_NAME struct {
	field fieldType
}

3. Method on the struct
func (x TYPE) method() returnType {
	return ...
}

4. Function that uses the interface
func doThing(v NAME) {
	v.method()
}

## Rule of thumb: use Sprintf only when the return type is string and you’re combining values into a sentence.

Learned
- Go implements interfaces and types implicitly
- Unlike other languages like Ts that uses the "implement" keyword, Go has no keyword for this. You just have to specify the methods in the interface.
- Go lets you write just the types (string, string) or the names and types (sourceFile string, destinationFile string). Both compile. The names are there for humans, so anyone reading the interface knows what goes where.
- interface Copier {
  copy(sourceFile: string, destinationFile: string): number
}

## Day 15 - 21/10/25
- Go Ch05 lesson
- Linux Ch06 lesson 

## Day 14 - 20/10/25
- Go Ch04 lesson 5 to 10
- Go Ch05 lesson 1 to 2
- Linux Ch06 lesson 1 to 3

Learned 
- I learnedd that memory layout in Go structs is very important.
- the order of fields in a struct can have a big impact on memory usagego
- You do not necessary have to worry about specifying the memory size but if it comes to size always lay it out from largest to the smallest.
- This order reduces memory usage.
- A string in a 64 bits machine is 16 bytes. This is because Go string does not hold the word itself. 
- string = [ pointer to the text: 8 bytes ][ length: 8 bytes ]
- int is 8 bytes, matches the machine word size
- int 32 is 4 bytes
- bool is 1 bytes
- Empty structs are Go's smallest possible type: they take up zero bytes of memory.


- Learned about interfaces in Go
- Interfaces let you write one function that works for many types, instead of rewriting it for each one.
- They help write reusable and scalable code that using defined behviours like methods that many types can share. this is because Go is type sensitive
- A normal interface has requirements and thats the specified method
- An empty interface however has no requirements, it accepts any type
- interface{}

## Day 13 - 19/10/25
- Go Ch04 - 1 - 5


Learned
- I learned about structs, you can have nested structs and embedded structs
- I also learned about anonymous structs, the ones that have no name and you just use once, like the one i used when writing tests. 
- I also learned that Go is not an OOP language like python






## Day 9 - 15/10/25
- Go - CLI Projects

Learned
- It’s a good habit: use %q for strings in test errors.

## Day 7 - 13/10/25
- Go - Cho4 - 12 -


Learned
- Learnt about anonymous functions in Go. Functions that have no name, if you're going to be using a function just once, why name it when you can just call it directly.
- This is something similar in Ts but different syntax. It is called arrow function in Ts and react.
- Examples in Ts and React are:
  nums.map((a) => a + a)
- How it is written in Go:

- Defer in Go. This means to "run this later, just before the function exits, no matter how it exits.
- You write it once at the top, and Go automatically runs it before any of the returns, including the “end of function” exit.
- This is the syntax : defer fmt.Println("3. deferred: runs last")
- If you have multiple defer statements in a single function, they are executed in last-in, first-out order (the last deferred call runs first).
- Something similar to this concept in React and Ts is: try and finally 
- Block scope in Go. This is similar to what is known in Ts, it is also block scope.
- Python however is function scoped.
- The analogy is that : Something you put in the cupboard can only be used inside the cupboard. Once you close the door. The door in this case is the curly bracket.






## Day 6 - 12/10/25
- Go - Ch04 - 5 - 11

Learned
- How to pass variables in functions
- When passing variales into functions, the funtion gets a copy and not the original.
- This is what the return statement does, you always have to return what you have specified that you will return
- When you declare parameters in afuntion, you must make use of all the parameter
- If you decide to want to ignore anyone, you have to do this with the underscore or the blank identifier
- Replace the unused variable with an underscore to not get any error message
- Learnt about named Returned Values
- You can just need a naked return instead of listing all that you have to return in the return statement
- A return statement without arguement is called a naked return
- Named return parameters are however good for documentation
- Naked returns should only be used in short and simple functions.

- Learnt about Nested and Guard clauses in Go functions
- Learnt about anonymous functions - this is a function without a name
- It looks like a normal function with the name removed
- func double(a int) int { return a + a }   // named
func       (a int) int { return a + a }   // anonymous, same thing without "double"
- When to use one: when the function is tiny and only used once, so it isn’t worth giving it a name.


## Day 5 - 11/10/25
- go - Ch04 - 1 - 4

Learned

- Functions in Go is very similar to TS because you have to specify the data type for the input paramater and for the return parameter, the only difference is that in Ts, you have to use a colon while in Go, you do no use any colon.

- Types in functions need to be stated explicitly after the parameter name
For example, func sub(x int, y int) int {
    return x-y
}



## Day 4 - 10/10/25
- Go - Ch03 - Lesson 5 to 9

Learned
- Learnt about Runes in Go
- Computers store texts as bytes. 
- Simple english letters need 1 byte
- Other characters like emojis might need more that 1 byte
- The way of storing text is UTF-8 and Go uses this
- A rune is Go name for one character
- To find the number of runes in a string = utf8.RuneCountInString(string)
- For example, the word "boot",has 4 runes
- Even emojis in Go counts as 1 rune


- Learnt about conditionals in Go
- It works the same way in Ts, just that you do not use parenthesis around the conditions
- Also, the curly braces has to be declared on the same line as the condition and not a new line
- if, else and else if are supported
- To make your code shorter and more concise, when declaring initial statements in conditionals, you declare it on a single line
- This is the syntax
- if length := getLength(email); length < 10 
- The semicolon ; separates “do this first” from “then check this condition”.

- Switch statement
- This is used when you want to compare a value against multiple options
- It is similar to the one in TS, unless that you do not need to add a break statement, Go knows when to break it self
- This is the syntax
switch plan {
case "basic":
	cost = 10.0
case "pro":
	cost = 20.0
default:
	cost = 0.0
}

- Wrote a function to calculate balanance using conditons


Struggled
- I had in deciding whether to use if and else if statments or switch statements
- 




## Day 3 - 09/10/25
- Go - Ch02 - Lesson 1 to 4

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


