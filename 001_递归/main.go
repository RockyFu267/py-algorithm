
package main

import (
	"fmt"
)

func main() {
	hanoi(3, "A", "B", "C")
}

func hanoi(n int64, a string, b string, c string) {

	if n > 0 { //检查是否还存在剩余的盘
		hanoi(n-1, a, c, b)
		fmt.Println(" moving from " + a + " to " + c)
		hanoi(n-1, b, a, c)
	}

}
