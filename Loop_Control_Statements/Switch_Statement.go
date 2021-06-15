package main

import "fmt"

func main() {
	var value string = "six"

	switch value {
	case "one":
		fmt.Println("OOP")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Golang")
	}
}
