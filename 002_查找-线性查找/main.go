package main

import "fmt"

func main() {
	var dataSet []int64
	var valueSet int64
	dataSet = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valueSet = 10
	linear_search(dataSet, valueSet)

}

func linear_search(dataSet []int64, valueSet int64) {
	for i := range dataSet {
		if dataSet[i] == valueSet {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("没有啊")
}
