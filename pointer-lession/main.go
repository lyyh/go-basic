package main

import "fmt"

func main() {
	// var a *int
	// a = new(int)
	// *a = 10
	// fmt.Println(*a)

	// var a *int = new(int)
	// *a = 10
	// fmt.Print(*a)

	var a map[string]int = make(map[string]int, 10)
	a["萨摩耶"] = 1
	fmt.Print(a)
}
