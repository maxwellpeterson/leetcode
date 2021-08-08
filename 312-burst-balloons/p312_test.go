package main

import "testing"

func TestOneElement(t *testing.T) {
	testCase(t, []int{5}, 5)
}

func TestTwoElementsUnordered(t *testing.T) {
	// Pop order doesn't matter here
	testCase(t, []int{1, 5}, 10)
}

func TestTwoElementsOrdered(t *testing.T) {
	// Pop order matters here
	testCase(t, []int{2, 5}, 15)
}

func TestLeetcodeExample(t *testing.T) {
	testCase(t, []int{3, 1, 5, 8}, 167)
}

func TestRepeating(t *testing.T) {
	testCase(t, []int{2, 2, 2, 2, 2}, 30)
}

func TestAllZeroes(t *testing.T) {
	testCase(t, []int{0, 0, 0}, 0)
}

func TestMixedZeroes(t *testing.T) {
	testCase(t, []int{0, 5, 0}, 5)
}

func testCase(t *testing.T, input []int, expected int) {
	output := maxCoins(input)
	if output != expected {
		t.Fatalf(`maxCoins(%v) = %v but want %v`, input, output, expected)
	}
}
