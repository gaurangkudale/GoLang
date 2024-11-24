package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to Slices")

	var fruitList = []string{"apple","Tomato","Peach"}
	
	fruitList = append(fruitList, "mango","Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 1
	highScore[1] = 11
	highScore[2] = 111
	highScore[3] = 1111

	fmt.Println(highScore)

	highScore = append(highScore, 22,332,3232)
	fmt.Println(len(highScore))
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)


	index1 := sort.Search(len(highScore), func(i int)bool  {
		return highScore[i] >= 1111
	})
	fmt.Println(index1)
	fmt.Println("---------------------------------------------------")

	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:4] // Create a slice from index 1 to 3 (exclusive)

	fmt.Println(mySlice) // Output: [2 3 4]

	mySlice[0] = 10
	fmt.Println(myArray) // Output: [1 10 3 4 5]


	//How to remove value frm slices based on index

	var courses = []string{"react.js","javascript","python","java","ruby"}
	fmt.Println(courses)

	ptr := &courses
	fmt.Println(ptr)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]... )
	fmt.Println(courses)
	

}