package main

func main() {}

func lengthOfLongestSubstring(input string) int {
	runes := []rune(input)

	if len(runes) < 2 {
		return len(runes)
	}

	currentLength := 0
	longestSubstring := 0
	nextIndex := map[rune]int{}

	for start := len(runes) - 1; start >= 0; start-- {
		next, exists := nextIndex[runes[start]]
		// If rune can be added to front of current substring without causing
		// duplicates, increase length of current substring.
		if !exists || start+currentLength < next {
			currentLength++
			longestSubstring = max(longestSubstring, currentLength)
		} else {
			// Otherwise, adding rune to front of current substring causes a
			// duplicate, so shorten current substring so the trailing duplicate
			// is not included.
			currentLength = next - start
		}

		nextIndex[runes[start]] = start
	}

	return longestSubstring
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
