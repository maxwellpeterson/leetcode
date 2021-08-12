package main

func main() {}

func averageWaitingTime(customers [][]int) float64 {
	totalWait := 0
	nextIdle := 0

	for _, pair := range customers {
		arrival, duration := pair[0], pair[1]
		preparationStart := max(nextIdle, arrival)

		totalWait += preparationStart - arrival + duration
		nextIdle = preparationStart + duration
	}

	return float64(totalWait) / float64(len(customers))
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
