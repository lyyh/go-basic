package main

import "fmt"

type student struct {
	name string
	age  int
}

func newPerson(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func (p student) Dream() {
	fmt.Printf("%v的梦想是学好go语言", p.name)
}

func (p *student) updateName() {
	p.name = "aaa"
}
func main() {
	p := newPerson("张三", 111)
	p.updateName()
	p.Dream()
}
