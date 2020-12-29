package main

import "testing"

func TestZeroByZero(t *testing.T) {
	testCase(t, [][]int{}, 0)
}

func TestOneByOne(t *testing.T) {
	testCase(t, [][]int{{5}}, 0)
}

func TestOneByTwo(t *testing.T) {
	testCase(t, [][]int{{5, 2}}, 0)
}

func TestTwoByOne(t *testing.T) {
	testCase(t, [][]int{{6}, {3}}, 0)
}

func TestTwoByTwo(t *testing.T) {
	testCase(t, [][]int{{1, 3}, {4, 2}}, 0)
}

func TestTwoByThree(t *testing.T) {
	testCase(t, [][]int{{1, 3, 5}, {4, 2, 7}}, 0)
}

func TestThreeByTwo(t *testing.T) {
	testCase(t, [][]int{{1, 3}, {4, 2}, {5, 11}}, 0)
}

func TestThreeByThreeTrap(t *testing.T) {
	testCase(t, [][]int{{1, 3, 2}, {4, 0, 2}, {0, 3, 5}}, 2)
}

func TestThreeByThreeDiamondTrap(t *testing.T) {
	testCase(t, [][]int{{0, 4, 1}, {5, 1, 4}, {2, 6, 0}}, 3)
}

func TestThreeByThreeSpillLeft(t *testing.T) {
	testCase(t, [][]int{{5, 4, 6}, {2, 2, 3}, {9, 4, 12}}, 0)
}

func TestThreeByThreeSpillRight(t *testing.T) {
	testCase(t, [][]int{{5, 5, 6}, {4, 3, 2}, {9, 6, 12}}, 0)
}

func TestThreeByThreeSpillUp(t *testing.T) {
	testCase(t, [][]int{{5, 1, 6}, {3, 0, 2}, {9, 4, 12}}, 1)
}

func TestThreeByThreeSpillDown(t *testing.T) {
	testCase(t, [][]int{{5, 4, 6}, {5, 1, 4}, {9, 3, 12}}, 2)
}

func TestSimpleElevation(t *testing.T) {
	testCase(t, [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}, 4)
}

func TestDifferentTrapElevations(t *testing.T) {
	testCase(t, [][]int{{1, 3, 3, 3, 2}, {2, 1, 4, 2, 4}, {3, 2, 2, 3, 1}}, 2)
}

func TestSpillUpLeft(t *testing.T) {
	testCase(t, [][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 1, 0}, {1, 0, 0, 0, 1}, {0, 1, 0, 1, 0}, {0, 0, 1, 0, 0}}, 0)
}

func TestComplexElevation(t *testing.T) {
	testCase(t, [][]int{{12, 13, 1, 12}, {13, 4, 13, 12}, {13, 8, 10, 12}, {12, 13, 12, 12}, {13, 13, 13, 13}}, 14)
}

func testCase(t *testing.T, input [][]int, expected int) {
	output := trapRainWater(input)
	if output != expected {
		t.Fatalf(`trapRainWater(%v) = %v but want %v`, input, output, expected)
	}
}
