package main

import "math"

func main() {}

// Running time: O(mn)
// Memory usage: O(mn)
func trapRainWater(height [][]int) int {
	if len(height) <= 2 || len(height[0]) <= 2 {
		return 0
	}

	rows, cols := len(height), len(height[0])

	// Create two m x n arrays for tracking the tallest columns up and to the
	// right from each column in the elevation. Also initialize the top and bottom
	// rows that don't get initialized by the loop below, ignoring their values
	// since these edge columns will never have any water on top of them
	tallestUp, tallestLeft := make([][]int, rows), make([][]int, rows)
	tallestUp[0], tallestLeft[0] = make([]int, cols), make([]int, cols)
	tallestUp[rows-1], tallestLeft[rows-1] = make([]int, cols), make([]int, cols)

	// Fill in the values for the tallestUp and tallestLeft arrays with a single
	// pass through the input array
	for row := 1; row < rows-1; row++ {
		tallestUp[row], tallestLeft[row] = make([]int, cols), make([]int, cols)
		for col := 1; col < cols-1; col++ {
			tallestUp[row][col] = max(tallestUp[row-1][col], height[row-1][col])
			tallestLeft[row][col] = max(tallestLeft[row][col-1], height[row][col-1])
		}
	}

	// On second pass through the input array, calculate the remaining tallest columns in
	// the down and right directions, which don't all need to be remembered. Use these
	// values, along with the tallestUp and tallestLeft arrays, to compute the maximum
	// water height for each cell and total water captured

	tallestDown := make([]int, cols)
	totalWater := 0

	for row := rows - 2; row >= 1; row-- {
		tallestRight := 0
		for col := cols - 2; col >= 1; col-- {
			tallestDown[col] = max(tallestDown[col], height[row+1][col])
			tallestRight = max(tallestRight, height[row][col+1])
			maxWaterHeight := min(tallestUp[row][col], tallestDown[col], tallestLeft[row][col], tallestRight)
			totalWater += max(0, maxWaterHeight-height[row][col])
		}
	}

	return totalWater
}

func min(vals ...int) int {
	minVal := math.MaxInt32
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
