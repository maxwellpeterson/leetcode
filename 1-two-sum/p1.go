package main

func main() {}

func twoSum(input []int, target int) []int {
	// Complements of elements of input array
	complements := map[int]int{}

	for index, value := range input {
		complementIndex, exists := complements[value]
		if exists {
			return []int{complementIndex, index}
		} else {
			complements[target-value] = index
		}
	}

	// Solution always exists
	return nil
}
