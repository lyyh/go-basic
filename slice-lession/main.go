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

// 在原始数组上进行过滤
func filter() []int {
	var a = []int{1, 2, 3, 4, 5}
	n := 0
	for _, v := range a {
		if v%2 == 0 {
			a[n] = v
			n++
		}
	}
	return a[:n]
}

func main() {
	// var a = make([]string, 5, 10)
	// fmt.Println(len(a), cap(a))
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)

	// sliceWithArr()

	// basic()

	// a := []int{1, 2, 3}
	// b := make([]int, len(a))
	// copy(b, a)
	// fmt.Println(b)
	// c := append([]int(nil), a...)
	// fmt.Println(c)
	// // 剪切将切片a中索引i~j位置的元素剪切掉。
	// d := append(a[:1], b[2:]...)
	// fmt.Println(d)

	// a := []int(nil)
	// a = append(a, 1, 2)
	// fmt.Printf("%#v", a)

	list := filter()

	fmt.Println("list", list)
}
