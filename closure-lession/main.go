package main

import (
	"fmt"
	"strings"
)

// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		return x + y
// 	}
// }

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	makeJpgFunc := makeSuffixFunc(".jpg")
	makeTxtFunc := makeSuffixFunc(".txt")
	fmt.Println(makeJpgFunc("test")) //test.jpg
	fmt.Println(makeTxtFunc("test")) //test.txt
	// var f = adder(10)
	// fmt.Println(f(10)) //10
	// fmt.Println(f(20)) //30
	// fmt.Println(f(30)) //60
}
