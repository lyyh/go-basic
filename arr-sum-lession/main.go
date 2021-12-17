package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArr(list [10]int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func main() {
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("%d", sum)
}
