package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是大佬"
	} else {
		talk = "你好"
	}
	return talk
}

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
