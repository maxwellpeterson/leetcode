package main

import "testing"

func TestEmptyString(t *testing.T) {
	testCase(t, "", 0)
}

func TestLonePair(t *testing.T) {
	testCase(t, "()", 2)
}

func TestPartialPair(t *testing.T) {
	testCase(t, "(()", 2)
}

func TestNestedPairs(t *testing.T) {
	testCase(t, "(()())", 6)
}

func TestPartialNestedPairs(t *testing.T) {
	testCase(t, ")()())", 4)
}

func TestLotsOfLeft(t *testing.T) {
	testCase(t, "()(((((()())((((()", 6)
}

func TestAllLeft(t *testing.T) {
	testCase(t, "((((", 0)
}

func TestLotsOfRight(t *testing.T) {
	testCase(t, "())))))(()(()())))))()((", 10)
}

func TestAllRight(t *testing.T) {
	testCase(t, "))))", 0)
}

func TestMultipleValid(t *testing.T) {
	testCase(t, "(()(()())((())", 8)
}

func TestFailedCase(t *testing.T) {
	testCase(t, ")()(((())))(", 10)
}

func testCase(t *testing.T, input string, expected int) {
	output := longestValidParentheses(input)
	if output != expected {
		t.Fatalf(`longestValidParentheses("%v") = %v but want %v`, input, output, expected)
	}
}
