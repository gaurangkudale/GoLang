package main

import (
	// "bufio"
	"fmt"
	"os"
	"strconv"
)

type Person struct{
	Name string
	Age int
	Country string
	Status bool
}

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
	s += sep + os.Args[i]
	sep = "hi"
	}
	fmt.Println(s)
	fmt.Println(os.Args[1:len(os.Args)])
	fmt.Println(isPalindrome(11211))
	// fmt.Println("Welcome to structs")
	// // no inheritance in glang and there is no super, parent 

	
	// var person Person
	// fmt.Println("Enter your name : ")
	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	person.Name = scanner.Text()
	// 	fmt.Println("You entered : ", person.Name)
	// }

	// gaurang := Person{person.Name,21, "India",true}
	// fmt.Println(gaurang)
	// fmt.Printf("%+v \n" ,gaurang)
	// gaurang.GetStatus()
}

//Methos 

// func (p Person) GetStatus(){
// fmt.Println("Is user activate : ",p.Status)
// }

// Paradrome Number
func isPalindrome(x int) bool {
    int_string := strconv.Itoa(x)
    n := len(int_string)
    for i := 0; i < len(int_string); i++ {
        if rune(int_string[i]) != rune(int_string[n - 1 - i]) {
            return false
        }
    }
    return true
}