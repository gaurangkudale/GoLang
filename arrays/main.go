package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in GoLang")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "mango"
	fruitList[3] = "peach"
	fmt.Println("Fruit List is : " , fruitList)
}