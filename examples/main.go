package main

import (
	"bufio"
	"fmt"
	"os"
)


type Person struct {
    name string
    age  int
}

func newPerson(name string) *Person {
    p := Person{name: name}
    p.age = 42
    return &p
}

func main() {
    mxIUNT := 4294967296
    fmt.Println(mxIUNT)
    fmt.Println(Person{"bob", 20})
    fmt.Println(newPerson("Gk"))
    fmt.Println(&Person{"gaurang", 23})
	a, j, z := 42, 2701, 0
	fmt.Println(a, j, z)
	fmt.Println(a)


	//Input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	txt, err := reader.ReadString('\n')
	
	fmt.Print("Your txt : ",txt, err)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
		
	}
}