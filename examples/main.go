package main

import (
    "fmt"
)

type maxInt uint64

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
    var mxIUNT maxInt = 4294967296
    fmt.Println(mxIUNT)

    fmt.Println(Person{"bob", 20})
    fmt.Println(newPerson("Gk"))
    fmt.Println(&Person{"gaurang", 23})
}