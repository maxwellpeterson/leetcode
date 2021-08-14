package main

func main() {}

func firstMissingPositive(input []int) int {
	validValue := func(value int) bool {
		return value > 0 && value <= len(input)
	}

	for ii, vv := range input {
		// If value is not in the correct position in the array.
		if validValue(vv) && vv != ii+1 {
			// If there is a collision with value that currently occupies the
			// correct position in the array, start swapping procedure.
			for validValue(input[vv-1]) && input[vv-1] != vv {
				vv, input[vv-1] = input[vv-1], vv
			}

			input[vv-1] = vv
		}
	}

	for ii, vv := range input {
		if vv != ii+1 {
			return ii + 1
		}
	}

	return len(input) + 1
}
