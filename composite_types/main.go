package main

import "fmt"

var array3ZeroValInts [3]int
var array3ValInts = [3]int{1: 5, 2}
var arrayCompDetSize = [...]int{1, 4: 60, 10: 12}

// When comparing arrays size of values
// Do not compare slices with ==, use selices.Equal(x,y)
// Create new slice
var x = []int{1, 2}

func main() {
	strMap := make(map[string]string)
	strMap["Japanese"] = "こんにちは世界"
	strMap["Arabic"] = "مرحبا بالعالم"
	strMap["Ukranian"] = "Привіт, світ"

	fmt.Println(strMap["Japanese"])
}
