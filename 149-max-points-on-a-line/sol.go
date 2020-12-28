package main

func main() {}

type point struct {
	x, y int
}

func makePoint(pos []int) point {
	return point{x: pos[0], y: pos[1]}
}

type slope struct {
	dy, dx int
}

func makeSlope(p0, p1 point) slope {
	if p0.x == p1.x {
		return slope{}
	}

	dy, dx := p1.y-p0.y, p1.x-p0.x
	gcd := gcd(dy, dx)
	dy, dx = dy/gcd, dx/gcd

	return slope{dy: dy, dx: dx}
}

// Running time: O(n^2)
// Memory usage: O(n^2) => Not sure about this, see how map is re-used in main loop
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	globalMaxCount := 0

	for ii := range points {
		// Tracks slope of line between current point and
		// every other not yet compared
		slopeCounts := map[slope]int{}
		// Tracks optimal value for current point
		localMaxCount := 1
		// Tracks number of duplicates of current point
		duplicateCount := 0

		// Compare current point to all others not yet compared
		for jj := ii + 1; jj < len(points); jj++ {
			p0, p1 := makePoint(points[ii]), makePoint(points[jj])

			// Check for duplicates and update counter if needed
			if p0 == p1 {
				duplicateCount++
				continue
			}

			// Calculate slope and update associated counter
			ss := makeSlope(p0, p1)
			if slopeCounts[ss] == 0 {
				slopeCounts[ss] = 2
			} else {
				slopeCounts[ss]++
			}

			// Update max slope count for this point if needed
			localMaxCount = max(localMaxCount, slopeCounts[ss])
		}

		// Update global max point count if needed
		globalMaxCount = max(globalMaxCount, localMaxCount+duplicateCount)
	}

	return globalMaxCount
}

// See en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(aa, bb int) int {
	for bb != 0 {
		aa, bb = bb, aa%bb
	}
	return aa
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
