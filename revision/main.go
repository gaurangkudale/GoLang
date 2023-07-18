package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Revision of pointers")
	// var ptr * int 
	// num := 32
	// ptr = &num
	// fmt.Println("Address if ptr is : ", ptr,"\nThe Value of num is : ", num)


	//Arrays 
	// a := [...]int{1,2:3,3:1,4,5,6,10:10}
	// fmt.Println(a)

	// var ab [2][3]int
	// fmt.Println(ab)
	// ab = [2][3]int{{1,2,3},{4,5,6}}
	// fmt.Println(ab)

	//Slieces
	// s := [] int {1,2,3,4,5}
	// s1 := s[2:5]
	// sCopy := make([]int, len(s1))
	// copy(sCopy, s1)
	// fmt.Println("Copy : ", sCopy)

	// fmt.Println("original slice is : ",s)
	// s = append(s, 6,7)
	// sa := make([]int, 4)
	// sa = append(sa)
	// fmt.Println("using make funtion in slice",sa)
	// fmt.Println("new slice is : ",s)
	// fmt.Println("Length of S is : " , len(s))
	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// 	fmt.Println("reverse slice is : ",s)
	// }

	// message := "Hello, world!"

	// // Iterating over a string using range
	// for index, char := range message {
	// 	fmt.Printf("Index: %d, Character: %c\n", index, char)
	// }
	// num := 1
	// var pointers *int = &num
	// fmt.Println(pointers)

	//Maps in GoLang
	myMap :=  make(map[string]int)
	myMap ["gk"] = 21
	fmt.Println(myMap["gk"])

	subjectMarks := map[string]float32 {"golang":1, "java":2, "python":3}
	fmt.Println(subjectMarks["golang"])
		
	student := make(map[string]int)

	student["raj"] = 1
	student["omkar"] = 2
	student["gaurang"] = 3

	for name,Value := range student{
		fmt.Println(name, Value)
	}
	fmt.Println(student)

}