package main

import (
	"fmt"
	"strings"
)

// func (b *Builder) String() string {
// 	return *(*string)(unsafe.Pointer(&b.buf))
// }

func main() {

	// str := "asong"
	// str = fmt.Sprintf("%s%s", str, str)
	// fmt.Println(str)
	baseSlice := []string{"asong", "真帅"}
	strings.Join(baseSlice, "")
	fmt.Println(baseSlice)
}
