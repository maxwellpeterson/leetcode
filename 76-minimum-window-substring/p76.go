package main

func main() {}

func minWindow(ss string, tt string) string {
	sequence, target := []rune(ss), []rune(tt)

	neededCount := map[rune]int{}
	totalNeeded := len(target)

	for _, rr := range target {
		neededCount[rr]++
	}

	start := 0
	optStart, optEnd := 0, len(sequence)-1

	for end, endRune := range sequence {
		if needed, exists := neededCount[endRune]; exists {
			if needed > 0 {
				totalNeeded--
			}

			neededCount[endRune]--

			if totalNeeded == 0 {
				for {
					if needed, exists := neededCount[sequence[start]]; exists && needed == 0 {
						if end-start < optEnd-optStart {
							optStart, optEnd = start, end
						}

						break
					} else if exists {
						neededCount[sequence[start]]++
					}

					start++
				}
			}
		}
	}

	if len(target) == 0 || totalNeeded > 0 {
		return ""
	} else {
		return string(sequence[optStart : optEnd+1])
	}
}
