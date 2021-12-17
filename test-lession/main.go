package main

import (
	"fmt"
	"math"
	"strings"
)

func TraversalString() {
	s := "aaa.博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Print(string(runeS2))
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func main() {
	TraversalString()
	changeString()
	sqrtDemo()
	// var a = "123adf"
	// fmt.Println(a)
	// fmt.Println(len(a))
	// strHanldered := fmt.Sprintf("%s%s", a, "asdf")
	// fmt.Println(strHanldered)
	var str []string = []string{"1", "2", "3"}
	s := strings.Join(str, "d")
	fmt.Println(s)
}
