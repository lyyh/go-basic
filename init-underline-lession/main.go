package main

import (
	"fmt"
	"os"
)

// import "basic/init-underline-lession/hello"

func main() {
	const (
		pi = 3.14
		e  = 2.71
	)
	const (
		a = 3
		b
		c
	)
	const (
		a1 = iota
		a2
		a3
	)
	fmt.Println(a, b, c)
	fmt.Println(a1, a2, a3)
	// hello.Say()
	buf := make([]byte, 1024)
	f, _ := os.Open("./a.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println(buf[:n])
		os.Stdout.Write(buf[:n])
	}
}
