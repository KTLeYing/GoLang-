package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type mystr string

type Student3 struct {
	Person
	int
	mystr
}

func main() {
	s1 := Student3{Person{"MZL", "man", 18}, 1, "kk"}
	fmt.Print(s1)
}
