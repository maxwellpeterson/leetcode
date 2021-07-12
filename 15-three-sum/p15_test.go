package main

import "testing"

func TestZero(t *testing.T) {
	testCase(t, []int{}, [][]int{})
}

func TestOne(t *testing.T) {
	testCase(t, []int{0}, [][]int{})
}

func TestTwo(t *testing.T) {
	testCase(t, []int{0, 0}, [][]int{})
}

func TestThreeValid(t *testing.T) {
	testCase(t, []int{0, 0, 0}, [][]int{{0, 0, 0}})
}

func TestThreeInvalid(t *testing.T) {
	testCase(t, []int{0, 1, 2}, [][]int{})
}

func TestExample(t *testing.T) {
	testCase(t, []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}})
}

func TestZeroes(t *testing.T) {
	testCase(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, [][]int{{0, 0, 0}})
}

func testCase(t *testing.T, input []int, expected [][]int) {
	output := threeSum(input)
	if !outputMatches(output, expected) {
		t.Fatalf(`threeSum(%v) = %v but want %v`, input, output, expected)
	}
}

func outputMatches(output, expected [][]int) bool {
	if len(output) != len(expected) {
		return false
	}

	for ii, aa := range output {
		if len(aa) != len(expected[ii]) {
			return false
		}

		for jj, bb := range aa {
			if bb != expected[ii][jj] {
				return false
			}
		}
	}
	return true
}
