package main

import "strings"

func countDistinctWords(messages []string) int {

	seen := make(map[string]struct{})

	for _, msg := range messages{

		for _, word := range strings.Fields(msg){
			lowerDistinctWord := strings.ToLower(word)

		seen[lowerDistinctWord] = struct{}{}


		}
		
		
		
		


	}

	return len(seen)
	
}
