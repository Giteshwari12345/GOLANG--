package main

import "fmt"

func main() {
	var x int = 0

	for x < 10 {
		if x == 4 {

			x = x + 4
			continue
		}
		fmt.Printf("value is: %d\n", x)
		x++
	}
}
