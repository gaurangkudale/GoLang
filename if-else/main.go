package main

import "fmt"

func main() {
	fmt.Println("WElcome to if-else")

	

	var result string

	if loginCount := 20; loginCount < 10 {
		result = "regular user"
	}else {
		result = "loginCount is different"
	}
	fmt.Println(result)
}
