package main

import "fmt"

type Person1 struct {
	name string
	sex  string
	age  int
}

type Student1 struct {
	*Person1
	id  int
	add string
}

func main() {
	s1 := Student1{&Person1{"MZL", "man", 18}, 1, "kk"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person1.name)
}
