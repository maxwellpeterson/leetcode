package main

func main() {}

func maxCoins(values []int) int {
	length := len(values)
	opt := make([][]int, length)

	localAcc := 0
	localMax := 0

	globalAcc := 0
	globalMax := 0

	for start := length - 1; start >= 0; start-- {
		opt[start] = make([]int, length)

		for end := start; end < length; end++ {
			localMax = 0

			for finalPop := start; finalPop <= end; finalPop++ {
				localAcc = 0

				if finalPop > start {
					localAcc += opt[start][finalPop-1]
				}

				localAcc += valueAt(values, start-1) * values[finalPop] * valueAt(values, end+1)

				if finalPop < end {
					localAcc += opt[finalPop+1][end]
				}

				localMax = max(localMax, localAcc)

			}

			opt[start][end] = localMax

			if start == 0 {
				globalAcc = 0

				if end > 0 {
					globalAcc += opt[0][end-1]
				}

				globalAcc += values[end]

				if end < length-1 {
					globalAcc += opt[end+1][length-1]
				}

				globalMax = max(globalMax, globalAcc)
			}
		}
	}

	return globalMax
}

func valueAt(values []int, index int) int {
	if index >= 0 && index < len(values) {
		return values[index]
	} else if index == -1 || index == len(values) {
		return 1
	} else {
		panic("Index too far out of bounds!")
	}
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
