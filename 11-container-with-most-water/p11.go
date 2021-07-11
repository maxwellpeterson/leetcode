package main

func main() {}

func maxArea(heights []int) int {
	ll := 0
	rr := len(heights) - 1

	maxAcc := 0

	for ll < rr {
		maxAcc = max(maxAcc, min(heights[ll], heights[rr])*(rr-ll))
		if heights[ll] < heights[rr] {
			ll++
		} else {
			rr--
		}
	}

	return maxAcc
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
