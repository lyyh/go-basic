package hello

import "fmt"

func init() {
	fmt.Println("我是subpackage")
}

func Say() {
	fmt.Println("say hello")
}
