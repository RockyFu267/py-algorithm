package main

import "fmt"

func main() {
	var listSet []int64
	listSet = []int64{9, 8, 7, 1, 2, 0, 3, 4, 6, 5}
	bubble_soft(listSet)
}

func bubble_soft(listSet []int64) {
	var j int
	var tmp int64
	for i := range listSet {
		for j <= len(listSet)-i-1 {
			if listSet[j] > listSet[j+1] {
				tmp = listSet[j+1]
				listSet[j+1] = listSet[i]
				listSet[i] = tmp
			}
			j = j + 1
		}
	}
	fmt.Println(listSet)
}
