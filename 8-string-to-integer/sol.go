package main

import "math"

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

		// Stop parsing if we hit a non-digit (this could
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

// There is probably a better way to do this without using 64-bit integers...
func checkOverflow(value, digit int, isPositive bool) (adjusted int, overflow bool) {
	// Safely compute new value as 64-bit integer
	newValue := 10*int64(value) + int64(digit)

	// Check for positive overflow
	if isPositive && newValue > int64(math.MaxInt32) {
		return math.MaxInt32, true
	}

	// Check for negative overflow
	if !isPositive && -newValue < int64(math.MinInt32) {
		return math.MinInt32, true
	}

	// Otherwise, no overflow
	return int(newValue), false
}
