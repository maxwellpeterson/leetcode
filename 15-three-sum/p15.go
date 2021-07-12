package main

import "sort"

func main() {}

func threeSum(nums []int) [][]int {
	results := [][]int{}

	sort.Ints(nums)

	for ii := 0; ii < len(nums); ii = shiftRight(ii, nums) {
		if len(nums[ii:]) >= 3 {
			jj := ii + 1
			kk := len(nums) - 1
			for jj < kk {
				if nums[ii]+nums[jj]+nums[kk] < 0 {
					jj++
				} else if nums[ii]+nums[jj]+nums[kk] == 0 {
					results = append(results, []int{nums[ii], nums[jj], nums[kk]})
					jj = shiftRight(jj, nums)
					kk = shiftLeft(kk, nums)
				} else {
					kk--
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
