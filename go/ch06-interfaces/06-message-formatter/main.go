package main

import (
	"fmt"

	
)

// the interface
type formatter interface{
	format() string
}

// the structs next

type plainText struct{
	message string
}
type bold struct{
	message string
}

type code struct{
	message string
}

//now function for each struct, this connets the method to the structs
func (p plainText) format() string{
	return fmt.Sprintf("%s", p.message)
}  

func (b bold) format() string{
	return fmt.Sprintf("**%s**", b.message)

}

func (c code) format() string{
	return fmt.Sprintf("`%s`", c.message)

}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
