package main

import "fmt"

func twoSum(list [10]int, target int) {
	for i, v := range list {
		other := target - v
		fmt.Println("other", other)
		for j := i + 1; j < len(list); j++ {
			if other == list[j] {
				fmt.Printf("(%d,%d)\n", v, other)
			}
		}
	}

}
func main() {
	b := [10]int{1, 7, 4, 2, 3, 4}
	twoSum(b, 8)
}
