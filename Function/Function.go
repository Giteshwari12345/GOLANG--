package main
import "fmt"

func area(length, width int)int{
	
Arr := length* width
	return Arr
}

func main() {
	

fmt.Printf("Area of rectangle is : %d", area(12, 10))
}
