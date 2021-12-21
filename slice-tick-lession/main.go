package main

import "fmt"

func filterOne(a []int, target int) []int {
	var n int = 0
	for idx, v := range a {
		if target != v {
			a[idx] = v
			n++
		}
	}
	return a[:n]
}

// 将切片a中索引i~j位置的元素剪切掉。
func cutFragment(sli []int, start int, end int) []int {
	// if start < 0 || end >= len
	return append(sli[:start], sli[end+1:]...)
}

func main() {
	var sli []int = []int{1, 2, 3, 4}
	sli = filterOne(sli, 2)
	fmt.Println("slice", sli)
}
