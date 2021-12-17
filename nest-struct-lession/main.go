package main

import "fmt"

// type Address struct {
// 	info string
// }
// type User struct {
// 	name string
// 	Address
// 	// address Address
// }

type Animal struct {
	name string
}

type Dog struct {
	Feet int8
	*Animal
}

func (a *Animal) move() {
	fmt.Printf("%v在跑\n", a.name)
}

func (d *Dog) eat() {
	fmt.Printf("%v在吃\n", d.Feet)
}

// 面向对象的继承
func main() {
	d := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "啊啊",
		},
	}

	d.move()
	d.eat()
}
