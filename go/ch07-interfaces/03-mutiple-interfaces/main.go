package main

import "fmt"

func (e email) cost() int {
	bodyLen := len(e.body)
	if !e.isSubscribed {
		
		return bodyLen * 5
	}
	return bodyLen * 2
	
}

func (e email) format() string {

	if !e.isSubscribed{
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
	return fmt.Sprintf("'%s' | Subscribed", e.body)
	
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
