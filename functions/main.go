package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in GoLang")
	// a := [...]int{1,2,4:2,3,6:-1}
	// fmt.Println(a)

	// s := make([]string,3)
	// s[0] = "hi"
	// s[1] = "gk"
	// s = append(s, "1","2","3")

	// fmt.Println(s)

	result := addition(5,4)
	fmt.Println(result)

	fmt.Println("Using PROAdder")
	fmt.Println(proAdder(1,2,3,4,5,6,7,8,9))
	
}

func proAdder (value ...int) int {
	total := 0

	for _, val := range value {
		total += val
	}
	return total
}

func addition (valOne int, valTwo int) int {
	return valOne + valTwo
}

func twoSum(nums []int, target int) []int {
    for i, v := range nums {
        for i = 0; i <= len(nums)-1; i++{
			fmt.Println(i)
		}
    }
    return nums[]
}