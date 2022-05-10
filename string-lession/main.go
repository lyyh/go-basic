package main

import (
	"fmt"
	"reflect"
)

// func (b *Builder) String() string {
// 	return *(*string)(unsafe.Pointer(&b.buf))
// }

type Student struct {
	name string
	age  int
}

func (stu *Student) study() {
	fmt.Printf("%s is studying", stu.name)
	stu.name = "啊啊啊"
}
func (stu *Student) backhome() {
	fmt.Printf("%s is backhome", stu.name)
}

func handleStu() {
	var stu = Student{
		"呃呃呃", 123,
	}
	stu.study()
	stu.backhome()
}

func handleInit() {
	var stu = new(Student)
	stu.name = "123"
	fmt.Printf("%#v", stu)
	var ret = fmt.Sprintf("%s %d", stu.name, stu.age)
	fmt.Println(ret)
}

func handleString() {
	// var str1 string = "Golang"
	var str2 string = "Go语言"

	var runeArr = []rune(str2) // 切片

	fmt.Println(reflect.TypeOf(runeArr[2]))
	// // []byte 类型处理
	// fmt.Printf("%d %c\n", str2[2], str2[2]) // 232 è 一个字节存不下
	// fmt.Println("len(str2)：", len(str2))    // len(str2)： 8
	// var runeStr = []rune(str2)
	// fmt.Printf("%c", runeStr[2])
}
func main() {
	handleStu()
	// var s = string(nil)
	// str := "asong"
	// str = fmt.Sprintf("%s%s", str, str)
	// fmt.Println(str)
	// baseSlice := []string{"asong", "真帅"}
	// strings.Join(baseSlice, "")
	// fmt.Println(baseSlice)

	// handleString()

	handleInit()
}
