package main

import "testing"

func TestEmptyString(t *testing.T) {
	testCase(t, "", 0)
}

func TestSingleCharacter(t *testing.T) {
	testCase(t, "a", 1)
}

func TestTwoCharacters(t *testing.T) {
	testCase(t, "ab", 2)
}

func TestThreeCharacters(t *testing.T) {
	testCase(t, "aba", 2)
}

func TestRepeatedCharacter(t *testing.T) {
	testCase(t, "aaaaaa", 1)
}

func TestDoubleRepeatedCharacter(t *testing.T) {
	testCase(t, "aaabbb", 2)
}

func TestRepeatedSubstring(t *testing.T) {
	testCase(t, "abcabcbb", 3)
}

func TestStrictSubstring(t *testing.T) {
	testCase(t, "pwwkew", 3) // Can't be a subsequence
}

func TestOffsetSubstring(t *testing.T) {
	testCase(t, "afghartw", 7) // Don't be greedy
}

func TestDoubleOffsetSubstring(t *testing.T) {
	testCase(t, "afghartgwyuoi", 10)
}

func TestSubstringTelescope(t *testing.T) {
	testCase(t, "abcabcdabcde", 5)
}

func testCase(t *testing.T, input string, expected int) {
	output := lengthOfLongestSubstring(input)
	if output != expected {
		t.Fatalf(`lengthOfLongestSubstring(%v) = %v but want %v`, input, output, expected)
	}
}
