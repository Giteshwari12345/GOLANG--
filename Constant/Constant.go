package main

import "fmt"

const PI = 3.14

func main() {
	const HLW = "HelloWorld"
	fmt.Println("Hello", HLW)

	fmt.Println("Happy", PI, "Day")

	const Correct = true
	fmt.Println("Go rule?", Correct)
}
