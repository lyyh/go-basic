package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type NewInt int  // 类型定义
type MyInt = int // 类型别名

type person struct {
	username, city string
	age            int
}

// 匿名结构体
func anonymity() {
	var p person
	p.age = 1
	p.username = "张三"
	p.city = "成都"
	fmt.Printf("%#v", p)

	var user struct {
		Name string
		Age  int
	}

	user.Name = "啊啊啊"
	user.Age = 18
}

// 结构体指针化
func pointerfly() {
	var p = new(person)
	p.username = "aaa"
	p.age = 123
	fmt.Printf("%#v", p)
}

// 通过取地址的操作来进行实例化
func instantiation() {
	var p = &person{
		username: "aaaa",
		age:      123,
	}
	fmt.Printf("%#v", p)
}

// 使用值的列表初始化
func instantiationDefault() {
	var p = &person{
		"aaaa",
		"aaaa",
		123,
	}
	fmt.Printf("%#v", p)
}

type student struct {
	name string
	age  int
}

func test() {
	stus := []student{
		{name: "aa", age: 11},
		{name: "bb", age: 22},
	}

	m := make(map[string]*student)

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	fmt.Printf("%#v", m)
}

// 构造函数
func newPerson(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}
func main() {
	// test()
	// instantiationDefault()
	// instantiation()
	// pointerfly()
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())

	// var a NewInt
	// var b MyInt

	// fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	// fmt.Printf("type of b:%T\n", b) //type of b:int

}
