package main

import (
	"bufio"
	"fmt"

	"os"

)

type Person struct{
	Name string
	Age int
	Country string
}

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in glang and there is no super, parent 

	
	var person Person
	fmt.Println("Enter your name : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		person.Name = scanner.Text()
		fmt.Println("You entered : ", person.Name)
	}

	gaurang := Person{person.Name,21, "India"}
	fmt.Println(gaurang)
	fmt.Printf("%+v \n" ,gaurang)

}