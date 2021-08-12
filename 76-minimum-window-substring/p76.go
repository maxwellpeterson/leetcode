package main

func main() {}

func minWindow(ss string, tt string) string {
	sequence, target := []rune(ss), []rune(tt)

	neededRunes := map[rune]int{}
	remainingNeededRunes := len(target)

	for _, rr := range target {
		neededRunes[rr]++
	}

	start, end := 0, 0
	shortestStart, shortestEnd := 0, len(sequence)

	for ; end < len(sequence); end++ {
		if needed, exists := neededRunes[sequence[end]]; exists {
			neededRunes[sequence[end]]--

			if needed > 0 {
				remainingNeededRunes--
			}
			if remainingNeededRunes == 0 {
				for {
					if needed, exists := neededRunes[sequence[start]]; exists && needed == 0 {
						if end+1-start < shortestEnd-shortestStart {
							shortestStart, shortestEnd = start, end+1
						}

						break
					} else if exists {
						neededRunes[sequence[start]]++
					}

					start++
				}
			}

		} else if remainingNeededRunes == len(target) {
			start++
		}
	}

	if len(target) == 0 || remainingNeededRunes > 0 {
		return ""
	} else {
		return string(sequence[shortestStart:shortestEnd])
	}
}
