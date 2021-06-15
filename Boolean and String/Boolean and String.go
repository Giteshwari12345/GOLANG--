package main

import "fmt"

func main() {

	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	result1 := str1 == str2

	fmt.Println(result1)
	fmt.Printf("The type of result1 is %T\n", result1)

	str := "GeeksforGeeks"

	fmt.Printf("Length of the string is: %d", len(str))

	fmt.Printf("\nString is: %s", str)

	fmt.Printf("\nType of str is: %T", str)

}
