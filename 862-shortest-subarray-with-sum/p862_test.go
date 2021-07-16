package main

import "testing"

func TestOne(t *testing.T) {
	testCase(t, []int{1}, 1, 1)
}

func TestNone(t *testing.T) {
	testCase(t, []int{1, 2}, 4, -1)
}

func TestSmall(t *testing.T) {
	testCase(t, []int{2, -1, 2}, 3, 3)
}

func TestLarge(t *testing.T) {
	testCase(t, []int{2, 0, 2, 3, 2, 1, 5, -2, 7, -10, 5}, 8, 3)
}

func TestNegatives(t *testing.T) {
	testCase(t, []int{-2, 3, 4, -5}, 7, 2)
}

func TestFailed(t *testing.T) {
	testCase(t, []int{-34, 37, 51, 3, -12, -50, 51, 100, -47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6}, 151, 2)
}

func testCase(t *testing.T, input []int, kk int, expected int) {
	output := shortestSubarray(input, kk)
	if output != expected {
		t.Fatalf(`shortestSubarray(%v, %v) = %v but want %v`, input, kk, output, expected)
	}
}
