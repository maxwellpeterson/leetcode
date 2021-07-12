package main

import "testing"

func TestSmall(t *testing.T) {
	testCase(t, [][]int{{0, 2}, {1, 3}}, 3)
}

func TestLarge(t *testing.T) {
	testCase(t, [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16)
}

func testCase(t *testing.T, input [][]int, expected int) {
	output := swimInWater(input)
	if output != expected {
		t.Fatalf(`swimInWater(%v) = %v but want %v`, input, output, expected)
	}
}
