package main

//A slice of slices is a grid: rows, and each row holds a list of values.
//[][]int
/*
grid := [][]int{
	{1, 2, 3},    // grid[0], row 0
	{4, 5, 6},    // grid[1], row 1
}

grid[1]       // [4 5 6], a whole row
grid[1][2]    // 6, row 1, column 2 */


/*
The comparison in TS
const grid: number[][] = [[1, 2, 3], [4, 5, 6]]
grid[1][2]   // 6*/

func createMatrix(rows, cols int) [][]int {

	
	grid := [][]int{}

	for i := range rows{
		row := []int{} //- this creates an empty slice first
		for j := range cols{
			row = append(row, i * j)
		}

		grid = append(grid, row)


	}

	return grid
	

	

	
}
