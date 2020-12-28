package main

func main() {}

func trap(height []int) int {
	length := len(height)

	if length <= 2 {
		return 0
	}

	// Find the tallest column to the left of each column
	// in the height array
	tallestToLeft := make([]int, length)
	tallestToLeft[0] = 0
	for ii := 1; ii < length; ii++ {
		tallestToLeft[ii] = max(tallestToLeft[ii-1], height[ii-1])
	}

	// Find tallest column to the right of each column
	// in the height array
	tallestToRight := make([]int, length)
	tallestToRight[length-1] = 0
	for ii := length - 2; ii >= 0; ii-- {
		tallestToRight[ii] = max(tallestToRight[ii+1], height[ii+1])
	}

	// Calculate the units of water stacked on top of each column,
	// using the heights of the tallest columns on either side, which
	// form a basin where water can collect
	totalWater := 0
	for ii, colHeight := range height {
		maxWaterHeight := min(tallestToLeft[ii], tallestToRight[ii])
		totalWater += max(0, maxWaterHeight-colHeight)
	}

	return totalWater
}

func min(aa, bb int) int {
	if aa < bb {
		return aa
	}
	return bb
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
