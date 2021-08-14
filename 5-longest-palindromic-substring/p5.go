package main

func main() {}

func longestPalindrome(input string) string {
	runes := []rune(input)
	isPalindrome := make([][]bool, len(runes))

	bestStart := 0
	bestEnd := 0

	for start := len(runes) - 1; start >= 0; start-- {
		isPalindrome[start] = make([]bool, len(runes))
		isPalindrome[start][start] = true

		for end := start; end < len(runes); end++ {
			if end-start == 1 {
				// For two-character strings, just make sure first and second
				// characters are the same.
				isPalindrome[start][end] = runes[start] == runes[end]
			} else if end-start > 1 {
				// For longer strings, make sure characters at either end are
				// the same, and that the substring in the middle is a palindrome.
				isPalindrome[start][end] = runes[start] == runes[end] && isPalindrome[start+1][end-1]
			}

			// After making assignment, check if we've found a new longest substring. Ties are
			// broken by earlier start index.
			if isPalindrome[start][end] && end-start >= bestEnd-bestStart {
				bestStart = start
				bestEnd = end
			}
		}
	}

	return string(runes[bestStart : bestEnd+1])
}
