package main

func main() {}

type point struct {
	x, y int
}

func makePoint(pos []int) point {
	return point{x: pos[0], y: pos[1]}
}

// A line defined by y = mx + b where m and b are stored in reduced
// fractional form, meaning m = nNum/mDen and b = bNum/bDen. Except
// if we're dealing with a vertical line, in which case the isVert
// and xVal fields are used
type line struct {
	isVert                       bool
	xVal, mNum, mDen, bNum, bDen int
}

func makeLine(p0, p1 point) line {
	if p0.x == p1.x {
		return line{isVert: true, xVal: p0.x}
	}

	mNum, mDen := p1.y-p0.y, p1.x-p0.x
	mGcd := gcd(mNum, mDen)
	mNum, mDen = mNum/mGcd, mDen/mGcd

	bNum, bDen := p0.y*mDen-p0.x*mNum, mDen
	bGcd := gcd(bNum, bDen)
	bNum, bDen = bNum/bGcd, bDen/bGcd

	return line{mNum: mNum, mDen: mDen, bNum: bNum, bDen: bDen}
}

// Running time:
// Memory usage:
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	// Keeps track of number of times each distinct point appears in input
	duplicates := map[point]int{}
	for _, pp := range points {
		duplicates[makePoint(pp)]++
	}

	// Keeps track of the set of distinct points that fall on each line
	pointSets := map[line]map[point]bool{}
	// Keeps track of the total number of points that fall on each line
	pointCounts := map[line]int{}
	maxPointCount := 0

	for ii, pp := range points {
		for jj := ii + 1; jj < len(points); jj++ {
			// Construct line defined by current pair of points
			p0, p1 := makePoint(pp), makePoint(points[jj])
			currentLine := makeLine(p0, p1)

			// Make sure point set for current line has been instantiated
			if pointSets[currentLine] == nil {
				pointSets[currentLine] = map[point]bool{}
			}

			// Update set of points that fall on current line and best
			// total line count if needed
			currentPointSet := pointSets[currentLine]
			if !currentPointSet[p0] {
				currentPointSet[p0] = true
				pointCounts[currentLine] += duplicates[p0]
			}
			if !currentPointSet[p1] {
				currentPointSet[p1] = true
				pointCounts[currentLine] += duplicates[p1]
			}
			if pointCounts[currentLine] > maxPointCount {
				maxPointCount = pointCounts[currentLine]
			}
		}
	}

	return maxPointCount
}

// See en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(aa, bb int) int {
	for bb != 0 {
		aa, bb = bb, aa%bb
	}
	return aa
}
