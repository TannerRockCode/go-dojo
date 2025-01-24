package main

import "fmt"

func main() {
	var i int = 20
	fmt.Printf("i: %d\n", i)
	var f float64 = float64(i)
	fmt.Printf("f: %f\n", f)

	const const_value = 20
	var value_int int = const_value
	var value_float float64 = const_value
	fmt.Printf("value int: %d\n", value_int)
	fmt.Printf("value float: %f\n", value_float)

}
