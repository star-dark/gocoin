package main

import "fmt"

type info struct {
	name string
	age  int
}

func (i info) sayHello() {
	fmt.Printf("내 이름은 %s 내 나이는 %d야", i.name, i.age)
}

func main() {
	introduce := info{"대형", 23}
	introduce.sayHello()
}
