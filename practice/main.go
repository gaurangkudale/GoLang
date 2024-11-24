package main

import (
    "fmt"
    "slices"
)

func main() {
    nums := []int{3, 1, 4, 1, 5}
    fmt.Println("Before sorting:", nums)

    slices.Sort(nums)
    fmt.Println("After sorting:", nums)
}
