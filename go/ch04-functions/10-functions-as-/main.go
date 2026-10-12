package main

import "fmt"

func reformat(message string, formatter func(string) string) string {

	firstMessage := formatter(message)
	secondMessage := formatter(firstMessage)
	thirdMessage := formatter(secondMessage)
	result := fmt.Sprintf("TEXTIO: %s", thirdMessage)
	return result

}
