package main

func getNameCounts(names []string) map[rune]map[string]int {

	result := make(map[rune]map[string]int)
	for _, name := range names{
		letter := []rune(name)[0]

		if _, ok := result[letter]; !ok{
			result[letter] = make(map[string]int)


		}
		result[letter][name]++
	}
	


	return result
}
