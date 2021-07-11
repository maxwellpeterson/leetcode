package main

import "testing"

func TestOnes(t *testing.T) {
	testCase(t, []int{1, 1}, 1)
}

func TestSimple(t *testing.T) {
	testCase(t, []int{1, 2, 1}, 2)
}

func TestMedium(t *testing.T) {
	testCase(t, []int{4, 3, 2, 1, 4}, 16)
}

func TestComplex(t *testing.T) {
	testCase(t, []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49)
}

func testCase(t *testing.T, input []int, expected int) {
	output := maxArea(input)
	if output != expected {
		t.Fatalf(`maxArea(%v) = %v but want %v`, input, output, expected)
	}
}
