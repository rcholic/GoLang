package main

import "fmt"

func findTarget(nums []int, target int) int {
	var result int = -1 //default position
	if len(nums) == 0 {
		return result
	}

	var start int = 0
	var end int = len(nums) - 1
	var mid int

	for start+1 < end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] == target {
		result = start
	}

	if nums[end] == target {
		result = end
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 8}
	var target int = 5

	fmt.Println(findTarget(nums, target))
}
