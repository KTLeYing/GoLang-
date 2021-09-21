package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	person Person
	id     int
	addr   string
	name   string
}

func main() {
	var s Student
	s.name = "mzl"
	fmt.Println(s)
	s.person.name = "jhd"
	fmt.Print(s)
}
