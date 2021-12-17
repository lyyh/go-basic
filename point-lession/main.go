package main

import "fmt"

func test(x *[2]int) {
	fmt.Printf("x: %p\n", x)
	x[1] = 1000
}

func isEmpty() {
	var p *string
	if p != nil {
		fmt.Println("不为空指针")
	} else {
		fmt.Println("为空指针")
	}
}

func makeMemo() {
	var s *string
	s = new(string)
	*s = "hello world"
	fmt.Printf("Type: %T,value: %v", s, s)
}

func main() {
	makeMemo()
	// var a [2]int = [2]int{1, 2}
	// fmt.Printf("a: %p\n", &a)
	// test(&a)
	// fmt.Println(a)
	// a := 10
	// b := &a
	// fmt.Printf("a:%d ptr:%p\n", a, b)
	// fmt.Printf("b:%p type:%T\n", b, b)

	// isEmpty()
}
