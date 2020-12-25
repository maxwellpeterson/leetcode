package main

import "testing"

func TestEmptyString(t *testing.T) {
	testCase(t, "", 0)
}

func TestOnlyWhitespaceShort(t *testing.T) {
	testCase(t, " ", 0)
}

func TestOnlyWhitespaceLong(t *testing.T) {
	testCase(t, "     ", 0)
}

func TestZero(t *testing.T) {
	testCase(t, "0", 0)
}

func TestMinusZero(t *testing.T) {
	testCase(t, "-0", 0)
}

func TestPlusZero(t *testing.T) {
	testCase(t, "+0", 0)
}

func TestLeadingZero(t *testing.T) {
	testCase(t, "055", 55)
}

func TestMinusLeadingZero(t *testing.T) {
	testCase(t, "-00032", -32)
}

func TestPlusLeadingZero(t *testing.T) {
	testCase(t, "+0067", 67)
}

func TestSimple(t *testing.T) {
	testCase(t, "42", 42)
}

func TestMinus(t *testing.T) {
	testCase(t, "-237", -237)
}

func TestPlus(t *testing.T) {
	testCase(t, "+585", 585)
}

func TestPadded(t *testing.T) {
	testCase(t, " 415", 415)
}

func TestPaddedMinus(t *testing.T) {
	testCase(t, "     -11", -11)
}

func TestPaddedPlus(t *testing.T) {
	testCase(t, "   +58", 58)
}

func TestWithWords(t *testing.T) {
	testCase(t, "4193 with words", 4193)
}

func TestMinusWithWords(t *testing.T) {
	testCase(t, "-41ggg", -41)
}

func TestPlusWithWords(t *testing.T) {
	testCase(t, "+5893     apple    ", 5893)
}

func TestPrefixWords(t *testing.T) {
	testCase(t, "words and 789", 0)
}

func TestMinusGap(t *testing.T) {
	testCase(t, " - 85", 0)
}

func TestPlusGap(t *testing.T) {
	testCase(t, "  +    346", 0)
}

func TestCommaNumber(t *testing.T) {
	testCase(t, " 855,902", 855)
}

func TestSpaceNumber(t *testing.T) {
	testCase(t, " -211 705 398", -211)
}

func TestTrailingMinus(t *testing.T) {
	testCase(t, " 24 - 37", 24)
}

func TestTrailingPlus(t *testing.T) {
	testCase(t, " -191+210  ", -191)
}

func TestNegativeOverflow(t *testing.T) {
	testCase(t, "-91283472332", -2147483648)
}

func TestPositiveOverflow(t *testing.T) {
	testCase(t, "3903567366", 2147483647)
}

func TestNegativeBound(t *testing.T) {
	testCase(t, "-2147483648", -2147483648)
}

func TestPositiveBound(t *testing.T) {
	testCase(t, "2147483647", 2147483647)
}

func TestAboveNegativeBound(t *testing.T) {
	testCase(t, "-2147483647", -2147483647)
}

func TestBelowPositiveBound(t *testing.T) {
	testCase(t, "2147483646", 2147483646)
}

func testCase(t *testing.T, input string, expected int) {
	output := myAtoi(input)
	if output != expected {
		t.Fatalf(`myAtoi(%v) = %v but want %v`, input, output, expected)
	}
}
