package main

func main() {}

// Running time: O(n)
// Memory usage: O(n)
func longestPalindrome(input string) string {
	// Convert input string to slice of runes to make
	// indexing easier
	runes := []rune(input)

	// This case will break indexing below, so handle explicitly
	if len(runes) == 1 {
		return input
	}

	// Find the number of consecutively repeated runes, starting with each rune
	// in the input. repeats[ii] = kk  means there is a kk-length substring of
	// the same rune starting at index ii of the input
	repeats := countRepeats(runes)

	// lengths[ii] is the length of the longest palindromic substring that
	// starts at index ii
	lengths := make([]int, len(runes), len(runes))
	// Final element must have substring length of one
	lengths[len(runes)-1] = 1

	for ii := len(runes) - 2; ii >= 0; ii-- {
		// If we can add the current rune to either end of the best postfix substring
		// to extend the palindrome, this will give us the biggest potenial increase in length
		trailingIndex := ii + lengths[ii+1] + 1
		if trailingIndex < len(runes) && runes[ii] == runes[trailingIndex] {
			lengths[ii] = lengths[ii+1] + 2
		} else if ii+2 < len(runes) && repeats[ii] == 1 && runes[ii] == runes[ii+2] {
			// Otherwise, we need to start a new palindrome starting from this rune
			lengths[ii] = 3
		} else {
			// Otherwise, our new palindrome can only come from repeated characters
			lengths[ii] = repeats[ii]
		}
	}

	return findLongest(runes, lengths)
}

func countRepeats(runes []rune) []int {
	repeats := make([]int, len(runes), len(runes))
	repeats[len(runes)-1] = 1

	for ii := len(runes) - 2; ii >= 0; ii-- {
		if runes[ii] == runes[ii+1] {
			repeats[ii] = repeats[ii+1] + 1
		} else {
			repeats[ii] = 1
		}
	}

	return repeats
}

func findLongest(runes []rune, lengths []int) string {
	// Slice values must be positive
	maxIndex := 0
	maxLength := 0

	for ii, ll := range lengths {
		if ll > maxLength {
			maxIndex = ii
			maxLength = ll
		}
	}

	return string(runes[maxIndex : maxIndex+maxLength])
}
