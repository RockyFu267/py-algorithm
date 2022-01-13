package main

import "fmt"

func main() {
	var dataSet []int
	var valueSet int
	dataSet = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valueSet = 0
	BinarySearch(dataSet, valueSet)
}

func BinarySearch(dataSet []int, valueSet int) {
	var left, right, mid int
	left = 0
	right = len(dataSet)

	for left <= right {
		mid = (right + left) / 2
		fmt.Println("mid:", mid, "midd的值：", dataSet[mid], "left:", left, "right:", right)
		switch {
		case dataSet[mid] < valueSet:
			if left == len(dataSet)-1 {       //防止越界
				fmt.Println("没有啊")
				return
			}
			left = mid + 1
		case dataSet[mid] > valueSet:
			if right == 0 {
				fmt.Println("没有啊")
				return
			}
			right = mid - 1
		default:
			fmt.Println(mid)
			return
		}
	}
}
