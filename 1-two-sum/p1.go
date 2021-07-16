package main

func main() {}

func twoSum(nums []int, target int) []int {
	comps := map[int]int{} // Complements
	for index, value := range nums {
		compIndex, exists := comps[value]
		if exists {
			return []int{compIndex, index}
		}
		comps[target-value] = index
	}
	return nil // Solution always exists
}
