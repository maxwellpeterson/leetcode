package main

import (
	"testing"
)

func TestConsecutiveStart(t *testing.T) {
	testCase(t, []int{2, 7, 11, 15}, 9, []int{0, 1})
}

func TestConsecutiveMiddle(t *testing.T) {
	testCase(t, []int{4, -9, 17, 5}, 8, []int{1, 2})
}

func TestConsecutiveEnd(t *testing.T) {
	testCase(t, []int{3, 2, 4}, 6, []int{1, 2})
}

func TestSpreadStart(t *testing.T) {
	testCase(t, []int{8, 2, 6, -10, 2}, 14, []int{0, 2})
}

func TestSpreadMiddle(t *testing.T) {
	testCase(t, []int{45, 5, -9, -10, 12}, -5, []int{1, 3})
}

func TestSpreadEnd(t *testing.T) {
	testCase(t, []int{35, -9, 11, 28, 700, -161}, -133, []int{3, 5})
}

func TestLone(t *testing.T) {
	testCase(t, []int{2, 1}, 3, []int{0, 1})
}

func TestLoneRepeat(t *testing.T) {
	testCase(t, []int{3, 3}, 6, []int{0, 1})
}

func TestRepeatSpread(t *testing.T) {
	testCase(t, []int{3, 4, -27, 9432, 43, -27, -2323, 56}, -54, []int{2, 5})
}

func testCase(t *testing.T, nums []int, target int, expected []int) {
	output := twoSum(nums, target)
	if !sliceEqual(output, expected) {
		t.Fatalf(`twoSum(%v, %v) = %v but want %v`, nums, target, output, expected)
	}
}

// Taken from stackoverflow.com/a/15312097
func sliceEqual(aa, bb []int) bool {

	// Slices must both be nil or both be not nill
	if (aa == nil) != (bb == nil) {
		return false
	}

	// Slices must be the same length
	if len(aa) != len(bb) {
		return false
	}

	// Slices must have the same elements in the same order
	for ii, vv := range aa {
		if vv != bb[ii] {
			return false
		}
	}

	return true
}
