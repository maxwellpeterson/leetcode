package main

import (
	"math"
	"testing"
)

func TestExampleOne(t *testing.T) {
	testCase(t, [][]int{{1, 2}, {2, 5}, {4, 3}}, 5.00000)
}

func TestExampleTwo(t *testing.T) {
	testCase(t, [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}, 3.25000)
}

func TestOneCustomer(t *testing.T) {
	testCase(t, [][]int{{5, 2}}, 2.00000)
}

func testCase(t *testing.T, input [][]int, expected float64) {
	output := averageWaitingTime(input)
	if math.Abs(output-expected) > 0.00001 {
		t.Fatalf(`minWindow(%v) = %v but want %v`, input, output, expected)
	}
}
