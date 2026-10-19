package main

import "fmt"

// support methods that can be defined on structs.
// A receiver is a special parameter that syntactically goes before the name of the function.

/**

type rect struct {
  width int
  height int
}

func (r rect) area() int {
	return r.width * r.height   // "r" = the struct, like "this" in Go = this.width
}

**/
type authenticationInfo struct {
	username string
	password string
}

// create the method below


func (a authenticationInfo) getBasicAuth()string{
	return fmt.Sprintf("Authorization: Basic %s:%s" , a.username, a.password)

}
