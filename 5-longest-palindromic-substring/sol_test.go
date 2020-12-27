package main

import "testing"

func TestOneCharacter(t *testing.T) {
	testCase(t, "a", "a")
}

func TestTwoCharacterRepeated(t *testing.T) {
	testCase(t, "cc", "cc")
}

func TestTwoCharacterNotRepeated(t *testing.T) {
	testCase(t, "ef", "e")
}

func TestTie(t *testing.T) {
	testCase(t, "babad", "bab")
}

func TestOverlap(t *testing.T) {
	testCase(t, "ababad", "ababa")
}

func TestTrickyOverlap(t *testing.T) {
	testCase(t, "abacabacaba", "abacabacaba")
}

func TestOverlapRepeatedMiddle(t *testing.T) {
	testCase(t, "abaabad", "abaaba")
}

func TestFullLength(t *testing.T) {
	testCase(t, "cbbc", "cbbc")
}

func TestRepeatedCharacter(t *testing.T) {
	testCase(t, "cbbd", "bb")
}

func TestRepeatedPrefixTie(t *testing.T) {
	testCase(t, "cccbc", "ccc")
}

func TestOptimalRepeatedPrefix(t *testing.T) {
	testCase(t, "aaaaga", "aaaa")
}

func TestNotOptimalRepeatedPrefix(t *testing.T) {
	testCase(t, "wwwppw", "wppw")
}

func TestRepeatedPostfixTie(t *testing.T) {
	testCase(t, "abbaaaa", "abba")
}

func TestOptimalRepeatedPostfix(t *testing.T) {
	testCase(t, "fghgffffff", "ffffff")
}

func TestNotOptimalRepeatedPostfix(t *testing.T) {
	testCase(t, "wkllkwwwww", "wkllkw")
}

func testCase(t *testing.T, input string, expected string) {
	output := longestPalindrome(input)
	if output != expected {
		t.Fatalf(`longestPalindrome(%v) = %v but want %v`, input, output, expected)
	}
}
