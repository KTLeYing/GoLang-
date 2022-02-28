package main

import "fmt"

type Person2 struct {
	name string
	sex  string
	age  int
}

type Student2 struct {
	person Person2
	id     int
	addr   string
	name   string
}

func main() {
	var s Student2
	s.name = "mzl"
	fmt.Println(s)
	s.person.name = "jhd"
	fmt.Print(s)
}
