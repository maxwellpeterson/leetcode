package main

func main() {}

type point struct {
	x, y int
}

// A line defined by y = mx + b where m and b
// are stored in franctional form
type line struct {
	isVert                       bool
	xVal, mNum, mDen, bNum, bDen int
}

func makeLine(x0, y0, x1, y1 int) line {
	mNum := y1 - y0
	mDen := x1 - x0
	if mDen == 0 {
		return line{isVert: true, xVal: x0}
	}
	mGcd := gcd(mNum, mDen)
	mNum, mDen = mNum/mGcd, mDen/mGcd

	bNum := y0*mDen - x0*mNum
	bDen := mDen
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

	pointsPerLine := map[line]map[point]bool{}
	bestLineCount := 0

	for ii, p0 := range points {
		for jj := ii + 1; jj < len(points); jj++ {
			// Create line defined by current pair of points, and find recorded
			// points that fall on this line
			x0, y0, x1, y1 := p0[0], p0[1], points[jj][0], points[jj][1]
			currentLine := makeLine(x0, y0, x1, y1)

			// Second-tier map may not be instantiated yet
			if pointsPerLine[currentLine] == nil {
				pointsPerLine[currentLine] = map[point]bool{}
			}

			// Add current points to current line (map entries may already exist)
			pointsPerLine[currentLine][point{x0, y0}] = true
			pointsPerLine[currentLine][point{x1, y1}] = true

			if len(pointsPerLine[currentLine]) > bestLineCount {
				bestLineCount = len(pointsPerLine[currentLine])
			}
		}
	}

	return bestLineCount
}

// See en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(aa, bb int) int {
	for bb != 0 {
		aa, bb = bb, aa%bb
	}
	return aa
}
