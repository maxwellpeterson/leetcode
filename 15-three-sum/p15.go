package main

import "sort"

func main() {}

func threeSum(nums []int) [][]int {
	results := [][]int{}

	sort.Ints(nums)

	for first := 0; first <= len(nums)-3; first = shiftRight(first, nums) {
		if nums[first] <= 0 {
			middle := first + 1
			last := len(nums) - 1
			for middle < last && nums[last] >= 0 {
				if nums[first]+nums[middle]+nums[last] < 0 {
					middle++
				} else if nums[first]+nums[middle]+nums[last] == 0 {
					results = append(results, []int{nums[first], nums[middle], nums[last]})
					middle = shiftRight(middle, nums)
					last = shiftLeft(last, nums)
				} else {
					last--
				}
			}
		}
	}

	return results
}

func shiftLeft(index int, nums []int) int {
	startValue := nums[index]

	for index > 0 && nums[index] == startValue {
		index--
	}

	return index
}

func shiftRight(index int, nums []int) int {
	startValue := nums[index]

	for index < len(nums) && nums[index] == startValue {
		index++
	}

	return index
}
