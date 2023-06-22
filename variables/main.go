package main

import (
	"bufio"
	"fmt"
	"os"

	// "github.com/go-delve/delve/pkg/dwarf/reader"
)

func main() {
	// declaring variables
	// var username string = "Gaurang"
	// fmt.Println(username)
	// fmt.Printf("The username is : %s \n" , username)
	// fmt.Printf("variable is of type : %T  \n", username )

	// var smallVal int = 255
	// fmt.Println(smallVal)

	// numberOfUser := 545.33
	// fmt.Println(numberOfUser)

	// const name, age = "Kim", 22
	// fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

	// User Input

	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for pizza : ")
	
	//comma ok syntax OR error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}
