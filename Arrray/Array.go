package main

import "fmt"

func main() {

	var myarr [2]string

	myarr[0] = "Hello"
	myarr[1] = "GO-CODE"

	fmt.Println("Elements of Array:")
	fmt.Println("Element 1: ", myarr[0])
	fmt.Println("Element 2: ", myarr[1])

	arr := [4]string{"Hie", "Good Morning", "I am golang developer"}

	fmt.Println("\nElements of the array:")

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])

	}

}
