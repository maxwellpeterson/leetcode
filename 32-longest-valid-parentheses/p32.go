package main

func main() {}

type Pair struct {
	startIndex, maxLenBeforeStart int
}

// Running time: O(n)
// Memory usage: O(n) (But not GC efficient?)
func longestValidParentheses(input string) int {
	runes := []rune(input)

	maxLen := 0
	currentLen := 0

	stack := []Pair{}

	for ii, rr := range runes {
		if rr == '(' {
			stack = append(stack, Pair{ii, currentLen})
			currentLen = 0
		} else if len(stack) > 0 {
			ll := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Popped element can't get garbage collected?

			currentLen = ii - ll.startIndex + 1 + ll.maxLenBeforeStart

			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 0
		}
	}

	return maxLen
}
