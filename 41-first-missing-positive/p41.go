package main

func main() {}

func firstMissingPositive(input []int) int {
	currentMax := 1

	for ii, vv := range input {
		if vv <= 0 || vv > len(input) {
			input[ii] = -1
		} else {
			currentMax = max(currentMax, vv)
		}
	}

	for _, vv := range input {
		if vv > 0 && input[vv-1] > 0 { // "Collision"
			to := vv - 1 // Index

			// Start swapping procedure
			for input[to] > 0 {
				nextTo := input[to] - 1
				input[to] = 0
				to = nextTo
			}

			input[to] = 0
		} else if vv > 0 {
			input[vv-1] = 0
		}
	}

	for ii, vv := range input {
		if vv != 0 {
			return ii + 1
		}
	}

	return currentMax + 1
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
