package main

import "fmt"

func main() {
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var (
		i, j = 0, 1
	)
	for ; i < j && j < len(nums); {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}

	return i + 1
}
