package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to TIME package")

	presentTime := time.Now()
	fmt.Println(presentTime)

	formatedTime := presentTime.Format("01-02-2006")
	fmt.Println(formatedTime)

	createdDate := time.Date(2023,time.July, 01, 12,00,00,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday "))
}
