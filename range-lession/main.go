package main

import (
	"fmt"
	"reflect"
)

func main() {
	testString := "Hello，世界"

	for _, c := range testString {
		fmt.Printf("%c 的类型是 %s\n", c, reflect.TypeOf(c))
	}
}
