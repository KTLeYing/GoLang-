package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type mystr string

type Student struct {
	Person
	int
	mystr
}

func main() {
	s1 := Student{Person{"MZL", "man", 18}, 1, "kk"}
	fmt.Print(s1)
}
