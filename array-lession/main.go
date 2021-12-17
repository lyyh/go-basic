package main

import "fmt"

func TraversalString() {
	// var arrStr = make([]string, 5, 10)
	// fmt.Println("%v", arrStr)
}

// 全局
var arr0 [3]int = [3]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4}
var arr3 = [5]int{4: 1}

func main() {
	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)

	TraversalString()

	// 局部
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3, 4} //通过初始化确定数组长度
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 123},
		{"user2", 234},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)

	a1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b1 := [...][2]int{{1, 2}, {3, 4}}
	fmt.Println(a1)
	fmt.Println(b1)
	// d :=
	// b :=

	// 多维数组遍历
	f := [2][3]int{{2, 3, 4}, {1, 2, 3}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d) = %d", k1, k2, v2)
		}
		fmt.Println()
	}
}
