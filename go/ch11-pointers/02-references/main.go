package main

import (
	"strings"
)

func removeProfanity(message *string) {
	badWords := []string{"fubb", "shiz", "witch"}

	msg := *message

	// this makes a copy of the original message without touching the originalll
	// this loop below makes changes to the referecnece msg while the original stays intact
	for _, word := range badWords{
		stars := strings.Repeat("*", len(word))
		msg = strings.ReplaceAll(msg, word, stars)
	}

	// this line changes the original to take in the reference message

	*message = msg


	
}
