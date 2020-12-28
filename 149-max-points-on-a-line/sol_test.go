package main

import "testing"

func TestZeroPoints(t *testing.T) {
	testCase(t, [][]int{}, 0)
}

func TestOnePoint(t *testing.T) {
	testCase(t, [][]int{{5, 17}}, 1)
}

func TestTwoPoints(t *testing.T) {
	testCase(t, [][]int{{2, 8}, {5, -10}}, 2)
}

func TestThreePointsInLine(t *testing.T) {
	testCase(t, [][]int{{1, 1}, {2, 2}, {3, 3}}, 3)
}

func TestDuplicate(t *testing.T) {
	testCase(t, [][]int{{1, 1}, {2, 2}, {3, 3}, {3, 3}}, 4)
}

func TestCrossLines(t *testing.T) {
	testCase(t, [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4)
}

func TestCrossOrigin(t *testing.T) {
	testCase(t, [][]int{{-1, -1}, {0, 0}, {1, 1}}, 3)
}

func TestCrossOriginReverse(t *testing.T) {
	testCase(t, [][]int{{1, -1}, {0, 0}, {-1, 1}, {4, -4}, {-10, 10}}, 5)
}

func TestDoubleCrossOrigin(t *testing.T) {
	testCase(t, [][]int{{-1, -1}, {4, 4}, {1, 1}, {-2, 2}, {-1, 1}, {3, -3}, {10, -10}}, 4)
}

func TestVerticalLine(t *testing.T) {
	testCase(t, [][]int{{0, 1}, {0, 2}, {0, 3}}, 3)
}

func testCase(t *testing.T, input [][]int, expected int) {
	output := maxPoints(input)
	if output != expected {
		t.Fatalf(`maxPoints(%v) = %v but want %v`, input, output, expected)
	}
}
