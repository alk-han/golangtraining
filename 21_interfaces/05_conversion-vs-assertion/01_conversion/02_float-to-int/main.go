package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	// y gets narrower
	fmt.Println(int(y) + x) // 24
	// conversion: float64 to int
}
