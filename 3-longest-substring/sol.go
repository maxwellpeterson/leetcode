package main

func main() {}

// Running time: O(n)
// Memory usage: O(n)
func lengthOfLongestSubstring(input string) int {
	// Convert input string to slice of runes to make
	// indexing easier
	runes := []rune(input)

	// These base cases will break the indexing below, so handle them explicitly
	if len(runes) == 0 {
		return 0
	}
	if len(runes) == 1 {
		return 1
	}

	// Find the index of the next occurence of each rune in the input string,
	// which is len(input) if rune does not appear again
	nextIndex := findNextIndex(runes)

	// lengths[ii] is the length of the longest non-repeating substring that
	// starts at index ii
	lengths := make([]int, len(runes), len(runes))
	// Final element must have substring length of one
	lengths[len(runes)-1] = 1

	for ii := len(runes) - 2; ii >= 0; ii-- {
		// If rune can be added to front of postfix substring without causing
		// duplicates, substring starting with that rune is one longer than postfix
		if ii+lengths[ii+1] < nextIndex[ii] {
			lengths[ii] = lengths[ii+1] + 1
		} else {
			// Otherwise, adding rune to front of postfix substring causes a duplicate,
			// so shorten postfix such that trailing duplicate is not included
			lengths[ii] = nextIndex[ii] - ii
		}
	}

	// After computing the values for each start index, we still need to find the best one
	return findMaxValue(lengths)
}

func findNextIndex(runes []rune) []int {
	nextIndex := map[rune]int{}
	result := make([]int, len(runes), len(runes))

	for ii := len(runes) - 1; ii >= 0; ii-- {
		rr := runes[ii]
		next, exists := nextIndex[rr]
		// If rune appears later in the slice, set next index to be that index,
		// otherwise set it to the length of the slice
		if exists {
			result[ii] = next
		} else {
			result[ii] = len(runes)
		}
		// Update recorded index of rune
		nextIndex[rr] = ii
	}

	return result
}

func findMaxValue(ints []int) int {
	// For this problem, we know all of the slice values are positive
	max := 0

	for _, vv := range ints {
		if vv > max {
			max = vv
		}
	}
	return max
}
