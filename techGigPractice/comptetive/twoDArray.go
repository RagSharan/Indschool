package comptetive

import "fmt"

func readTwoDArrayFromConsole() {
	//read two D array from console
	var row int
	var col int

	fmt.Scan(&row, &col)
	var twoD = make([][]int, row)

	for i := 0; i < row; i++ {
		twoD[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Scan(&twoD[i][j])
		}
	}
	fmt.Print(twoD)
}

func readfromConsole() {
	//read two D array from console
	var row int
	var col int

	fmt.Scan(&row, &col)
	var twoD = make([][]int, row)

	var result = make([]int, 0)
	for i := 0; i < row; i++ {
		twoD[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Scan(&twoD[i][j])
			if twoD[i][j]%2 == 0 {
				result = append(result, twoD[i][j])
			}
		}
	}
	fmt.Print(result)
}

func readMatrixfromConsole() {
	//read two D array from console
	var row, col int

	fmt.Scan(&row, &col)
	var matrix = make([][]int, row)

	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	findDaigonal(matrix, row, col)
}

func findDaigonal(matrix [][]int, row int, col int) {
	length := row + col

	for i := 0; i < length; i++ {
		start_col := max(0, i-row)
		count := min(min(i, (col-start_col)), row)

		for j := 0; j < count; j++ {
			fmt.Print(matrix[min(row, i)-j-1][start_col+j], " ")
		}
		fmt.Println()
	}

}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
