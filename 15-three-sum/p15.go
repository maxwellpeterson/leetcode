package main

import (
	"sort"
	"sync"
)

func main() {}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	triples := make(chan []int)
	var waitGroup sync.WaitGroup

	findTriples := func(first int) {
		defer waitGroup.Done()

		middle := first + 1
		last := len(nums) - 1

		for middle < last && nums[last] >= 0 {
			if nums[first]+nums[middle]+nums[last] < 0 {
				middle++
			} else if nums[first]+nums[middle]+nums[last] == 0 {
				triples <- []int{nums[first], nums[middle], nums[last]}
				middle = shiftRight(middle, nums)
				last = shiftLeft(last, nums)
			} else {
				last--
			}
		}
	}

	go func() {
		for first := 0; first <= len(nums)-3; first = shiftRight(first, nums) {
			if nums[first] <= 0 {
				waitGroup.Add(1)
				go findTriples(first)
			}
		}

		waitGroup.Wait()
		close(triples)
	}()

	results := [][]int{}

	for triple := range triples {
		results = append(results, triple)
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
