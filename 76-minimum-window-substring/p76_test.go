package main

import "testing"

func TestEmptyT(t *testing.T) {
	testCase(t, "a", "", "")
}

func TestNoSolution(t *testing.T) {
	testCase(t, "a", "aa", "")
}

func TestFullMatch(t *testing.T) {
	testCase(t, "a", "a", "a")
}
func TestConsecutiveMatchOrdered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "jak", "jak")
}

func TestConsecutiveDuplicateMatchOrdered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "kaa", "kaa")
}

func TestConsecutiveMatchUnordered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "kab", "bka")
}

func TestConsecutiveDuplicateMatchUnordered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "kblb", "klbb")
}

func TestNonConsecutiveMatchOrdered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "abj", "abfej")
}

func TestNonConsecutiveDuplicateMatchOrdered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "jbb", "jaklbb")
}

func TestNonConsecutiveMatchUnordered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "abd", "bkaalfjd")
}

func TestNonConsecutiveDuplicateMatchUnordered(t *testing.T) {
	testCase(t, "abfejaklbbkaalfjd", "jkj", "jaklbbkaalfj")
}

func TestStart(t *testing.T) {
	testCase(t, "abdebcca", "ba", "ab")
}

func TestEnd(t *testing.T) {
	testCase(t, "abdebcca", "ca", "ca")
}

func testCase(t *testing.T, ss string, tt string, expected string) {
	output := minWindow(ss, tt)
	if output != expected {
		t.Fatalf(`minWindow(%v, %v) = %v but want %v`, ss, tt, output, expected)
	}
}
