package main

import (
	"fmt"
)

func main() {
	var result = recur(2, 5)
	fmt.Printf("Result %d\n", result)
}

func recur(num int, times int) int {
	if times <= 1 {
		return num
	}

	return num + recur(num, times-1)
}
