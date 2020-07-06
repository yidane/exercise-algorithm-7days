package main

import "fmt"

func main() {
	var arr = []int{1, 8, 6, 100, 100, 0, 8, 3, 1}
	fmt.Println(maxAreaStupid(arr))
	fmt.Println(maxArea(arr))
}

func maxAreaStupid(height []int) int {
	var maxSize = 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			var (
				leftHeight  = height[i]
				rightHeight = height[j]
				sizeHeight  = 0
				sizeWidth   = j - i
			)

			if leftHeight > rightHeight {
				sizeHeight = rightHeight
			} else {
				sizeHeight = leftHeight
			}

			var size = sizeHeight * sizeWidth

			if size > maxSize {
				maxSize = size
			}
		}
	}

	return maxSize
}

func maxArea(height []int) int {
	var (
		maxSize = 0
		i       = 0
		j       = len(height) - 1
	)

	for i < j {
		var (
			leftHeight  = height[i]
			rightHeight = height[j]
			curHeight   = 0
			width       = j - i
		)

		if leftHeight > rightHeight {
			curHeight = rightHeight
			j--
		} else {
			curHeight = leftHeight
			i++
		}

		var curSize = width * curHeight
		if curSize > maxSize {
			maxSize = curSize
		}
	}

	return maxSize
}
