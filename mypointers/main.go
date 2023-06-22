package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var mypointer *int
	// fmt.Println("Value of pointer is : ",mypointer)

	myNumber := 23 

	var ptr = &myNumber
	fmt.Println("value of actual prointer is : ", ptr)
	fmt.Println("value of actual prointer is : ", *ptr)

	newVal := *ptr + 2
	val := &newVal
	fmt.Println("NewVal : " , val)
	fmt.Println("NewVal : " , *val)
	fmt.Println("New value is : ",myNumber)
	fmt.Println("New value is : ",ptr)

	myValue := 50
	var myValuePtr = &myValue
	fmt.Println("Pointer to Myvalue is : ", myValuePtr)
	fmt.Println("Real value of Myvalue is : ", *myValuePtr)

	
	*myValuePtr = 100
	fmt.Println(myValuePtr)
	


}