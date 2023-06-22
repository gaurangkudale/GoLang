package main

import (
	"fmt"

	// "github.com/gogo/protobuf/test/indeximport-issue72/index"
)

func main() {
	fmt.Println("Welcome to Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday","Wednesday","Thusrsday","Friday","Saturday"}
	fmt.Println(days)

	//For Loop type 1

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// For loop type 2

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// For Loop type 3 
	
	// for index, day := range days{
	// 	fmt.Printf("For index %v the value is %v \n", index, day)
	// }

	
}