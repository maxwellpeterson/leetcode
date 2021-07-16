package main

import "testing"

func TestMissingLastElement(t *testing.T) {
	testCase(t, []int{1, 2, 0}, 3)
}

func TestMissingInMiddle(t *testing.T) {
	testCase(t, []int{3, 4, -1, 1}, 2)
}

func TestMissingOne(t *testing.T) {
	testCase(t, []int{7, 8, 9, 11, 12}, 1)
}

func TestMissingAtEndShort(t *testing.T) {
	testCase(t, []int{1}, 2)
}

func TestMissingAtEndLong(t *testing.T) {
	testCase(t, []int{5, 3, 1, 2, 4}, 6)
}

func TestMissingRepeated(t *testing.T) {
	testCase(t, []int{1, 2, 3, 3, 3, 2, 2, 1, 1, 1, 2, 3}, 4)
}

func TestMissingAllZero(t *testing.T) {
	testCase(t, []int{0, 0, 0, 0, 0, 0}, 1)
}

func TestMissingAllNegative(t *testing.T) {
	testCase(t, []int{-10, -2, -20, -680}, 1)
}

func TestMissingAllNonPositive(t *testing.T) {
	testCase(t, []int{0, -5, -17, 0, 0, 0, -8, -101}, 1)
}

func TestFailed(t *testing.T) {
	testCase(t, []int{-1, 4, 2, 1, 9, 10}, 3)
}

func testCase(t *testing.T, input []int, expected int) {
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)
	output := firstMissingPositive(input)
	if output != expected {
		t.Fatalf(`firstMissingPositive(%v) = %v but want %v`, inputCopy, output, expected)
	}
}
