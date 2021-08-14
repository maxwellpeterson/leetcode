package main

func main() {}

func trap(height []int) int {
	totalWater := 0
	left, right := 0, len(height)-1
	tallestToLeft, tallestToRight := 0, 0

	for left <= right {
		if tallestToLeft < tallestToRight {
			tallestToLeft = max(tallestToLeft, height[left])
			totalWater += tallestToLeft - height[left]
			left++
		} else {
			tallestToRight = max(tallestToRight, height[right])
			totalWater += tallestToRight - height[right]
			right--
		}
	}

	return totalWater
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
