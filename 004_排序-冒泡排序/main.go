package main

func main() {
	var listSet []int64
	listSet = []int64{9, 8, 7, 1, 2, 0, 3, 4, 6, 5}
	bubble_soft(listSet)
}

func bubble_soft(listSet []int64) {
	for i := range listSet {
		println(i)
	}

}
