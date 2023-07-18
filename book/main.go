package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	// fmt.Println("Program Name : " , os.Args[0])
	// fmt.Println("Command Line arguments :")

	// for i, args := range os.Args[0:]{
	// 	fmt.Printf("%d: %s\n", i+1, args)
	// }

	// var s, sep string
	// for i := 0; i < len(os.Args); i++ {
	// 	s = s + sep + os.Args[i]
	// 	sep = "\n"
	// }
	// fmt.Println(s)


	for i, args := range os.Args[0:] {
		fmt.Println(i, " : ",args )
	}
	
	
}