package main

import "fmt"

func main() {

	arr := [7]string{"Hello", "I", "am", "Golang", "developer"}

	fmt.Println("Array:", arr)

	myslice := arr[1:6]

	fmt.Println("Slice:", myslice)

	fmt.Printf("Length of the slice: %d", len(myslice))

	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
