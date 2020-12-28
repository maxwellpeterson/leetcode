package main

import "testing"

func TestEmpty(t *testing.T) {
	testCase(t, []int{}, 0)
}

func TestOneElement(t *testing.T) {
	testCase(t, []int{5}, 0)
}

func TestTwoElements(t *testing.T) {
	testCase(t, []int{2, 7}, 0)
}

func TestSimpleTrap(t *testing.T) {
	testCase(t, []int{1, 0, 1}, 1)
}

func TestTallerOnLeft(t *testing.T) {
	testCase(t, []int{2, 0, 1}, 1)
}

func TestTallerOnRight(t *testing.T) {
	testCase(t, []int{2, 0, 7}, 2)
}

func TestSkinnyTrap(t *testing.T) {
	testCase(t, []int{7, 0, 7}, 7)
}

func TestPyramid(t *testing.T) {
	testCase(t, []int{2, 4, 5, 3, 0}, 0)
}

func TestInvertedPyramid(t *testing.T) {
	testCase(t, []int{5, 4, 2, 3, 6}, 6)
}

func TestManyZeroes(t *testing.T) {
	testCase(t, []int{0, 1, 0, 0, 4, 0, 2, 3, 0, 0, 0}, 6)
}

func TestAllZeroes(t *testing.T) {
	testCase(t, []int{0, 0, 0, 0, 0, 0, 0}, 0)
}

func TestLoneBlock(t *testing.T) {
	testCase(t, []int{0, 0, 0, 2, 0, 0, 0, 0}, 0)
}

func TestLoneDoubleBlock(t *testing.T) {
	testCase(t, []int{0, 0, 5, 2, 0, 0, 0, 0}, 0)
}

func TestSimpleElevation(t *testing.T) {
	testCase(t, []int{4, 2, 0, 3, 2, 5}, 9)
}

func TestComplexElevation(t *testing.T) {
	testCase(t, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
}

func testCase(t *testing.T, input []int, expected int) {
	output := trap(input)
	if output != expected {
		t.Fatalf(`trap(%v) = %v but want %v`, input, output, expected)
	}
}
