package main

import "math"

const minIntQuotient = math.MinInt32 / 10
const minIntRemainder = math.MinInt32 % 10
const maxIntQuotient = math.MaxInt32 / 10
const maxIntRemainder = math.MaxInt32 % 10

func main() {}

func myAtoi(input string) int {
	// Convert input string to slice of runes to make
	// indexing easier
	runes := []rune(input)

	for ii, rr := range runes {
		// Skip over any leading whitespace
		if rr == ' ' {
			continue
		}

		// If we reach a sign, start parsing numeric value,
		// but we first need to determine the sign
		if rr == '-' || rr == '+' {
			return parseValue(runes[ii+1:], rr == '+')
		}

		// If we reach a digit, start parsing numeric value,
		// knowing that value is implicitly positive
		if parseDigit(rr) >= 0 {
			return parseValue(runes[ii:], true)
		}

		break // Otherwise, we have an illegal leading character
	}

	// If we get here, the input string is either empty or
	// only whitespace characters (or we have an illegal
	// leading character)
	return 0
}

func parseValue(runes []rune, isPositive bool) int {
	value := 0

	for _, rr := range runes {
		digit := parseDigit(rr)

		// Stop parsing if we hit a non-digit (this could even
		// potentially be the first rune)
		if digit < 0 {
			break
		}

		// Check for potential overflow when adding in new digit
		adjusted, overflow := checkOverflow(value, digit, isPositive)

		// If adding new digit would cause overflow, return bounded value
		if overflow {
			return adjusted
		}

		// Otherwise, value can be safely updated
		value = adjusted
	}

	// Apply sign before returning parsed value
	if isPositive {
		return value
	}

	return -value
}

func parseDigit(rr rune) int {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for ii, dd := range digits {
		if dd == rr {
			return ii // Index of rune is integer value
		}
	}
	return -1 // Couldn't find a match
}

// There is probably a more elegant way to do this, but this should be pretty fast...
func checkOverflow(value, digit int, isPositive bool) (adjusted int, overflow bool) {
	// Check for potential negative overflow
	if !isPositive && (-value < minIntQuotient || (-value == minIntQuotient && -digit < minIntRemainder)) {
		return math.MinInt32, true
	}

	// Check for potential positive overflow
	if isPositive && (value > maxIntQuotient || (value == maxIntQuotient && digit > maxIntRemainder)) {
		return math.MaxInt32, true
	}

	// Otherwise, we can safely combine value with digit without overflow
	return 10*value + digit, false
}
