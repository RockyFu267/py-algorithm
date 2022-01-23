package main

import "fmt"

func main() {
	var listdata []int64
	listdata = []int64{9, 8, 7, 1, 2, 0, 3, 4, 6, 0, 5}
	select_sort(listdata)
}

func select_sort(listdata []int64) {
	var minNum, tmpNum int64
	for i := range listdata {
		minNum = listdata[i]
		for j := i + 1; j <= len(listdata)-1; j++ {
			if minNum > listdata[j] {
				tmpNum = listdata[j]
				listdata[j] = listdata[i]
				listdata[i] = tmpNum
			}
		}
		fmt.Println(listdata)
	}

}
