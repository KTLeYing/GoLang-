package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type mystr string

type Student struct {
	*Person
	id  int
	add string
}

func main() {
	s1 := Student{&Person{"MZL", "man", 18}, 1, "kk"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)
}
