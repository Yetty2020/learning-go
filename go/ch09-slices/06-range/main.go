package main

// range is a shortcut for looping over a slice, so you don’t need to write the i := 0; i < len(...); i++ part yourself.
//for INDEX, ELEMENT := range SLICE {
//fruits := []string{"apple", "banana", "grape"}
//The element is a copy of the value at that index.
//for i, fruit := range fruits {
///Go refuses to compile if you create a variable and never use it
//So if you only need one of the two, put _ (the “blank identifier”) in place of the other:
// The comparison to this in Ts is the ForEach keyword
//Go is always index first, then item. TS’s forEach is item first.
func indexOfFirstBadWord(msg []string, badWords []string) int {

	for i, word := range msg{
		// is the index and word is the item
		for _, bad := range badWords{
			// this is to loop through each badword in the slice
			if word == bad{
				return i
			}
		}
	}
	
	return -1
}
