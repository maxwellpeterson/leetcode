package main

func main() {}

func longestPalindrome(input string) string {
	runes := []rune(input)

	// isPalindrome[row][col] is true if the substring that starts at row and
	// ends at col is a palindrome, and false otherwise
	isPalindrome := make([][]bool, len(runes))
	for row := range isPalindrome {
		isPalindrome[row] = make([]bool, len(runes))
		isPalindrome[row][row] = true
	}

	bestRow := 0
	bestCol := 0

	for row := len(runes) - 1; row >= 0; row-- {
		for col := row; col < len(runes); col++ {
			if col-row == 1 {
				// For two-character strings, just make sure first and second
				// characters are the same
				isPalindrome[row][col] = runes[row] == runes[col]
			} else if col-row > 1 {
				// For longer strings, make sure characters at either end are
				// the same, and that the substring in the middle is a palindrome
				isPalindrome[row][col] = runes[row] == runes[col] && isPalindrome[row+1][col-1]
			}

			// After making assignment, check if we've found a new longest substring. Ties are
			// broken by earlier start index
			if isPalindrome[row][col] && col-row >= bestCol-bestRow {
				bestRow = row
				bestCol = col
			}
		}
	}

	return string(runes[bestRow : bestCol+1])
}
