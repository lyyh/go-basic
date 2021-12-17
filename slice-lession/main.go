package main

import "fmt"

func basic() {
	var s1 []int
	if s1 == nil {
		fmt.Println("为空")
	} else {
		fmt.Println("不为空")
	}

	s2 := []int{}

	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	var s4 []int = make([]int, 0, 2)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	arr := [5]int{1, 2, 3, 4, 5}
	s6 := arr[1:4] // 前包后不包
	fmt.Println(s6)
}

func sliceWithArr() {
	data := make([]int, 2, 5)
	data = append(data, 0, 0, 0, 0)
	data = append(data, []int{1, 2, 3}...) // 追加一个切片需要解包
	fmt.Printf("%v\n", data)
}

func main() {
	// var a = make([]string, 5, 10)
	// fmt.Println(len(a), cap(a))
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)

	sliceWithArr()

	// basic()
}
