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
	// Value at zero index is zeroed by default
	for ii := 1; ii < length-1; ii++ {
		tallestToLeft[ii] = max(tallestToLeft[ii-1], height[ii-1])
	}

	// Find tallest column to the right of each column in the height array,
	// and accumulate the total units of water stacked on top of each column,
	// using the heights of the tallest columns on either side, which
	// form a basin where water can collect

	tallestToRight := 0
	totalWater := 0

	for ii := length - 2; ii >= 1; ii-- {
		tallestToRight = max(tallestToRight, height[ii+1])
		maxWaterHeight := min(tallestToLeft[ii], tallestToRight)
		totalWater += max(0, maxWaterHeight-height[ii])
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
