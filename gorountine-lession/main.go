package main

import (
	"fmt"
	"sync"
)

// func hello() {
// 	fmt.Print("hello goroutine done!")
// }

// func main() {
// 	go hello()
// 	fmt.Print("main goroutine done!")
// 	time.Sleep(time.Second)
// }

var wg sync.WaitGroup

func hello(num int) {
	defer wg.Done()
	fmt.Println("hello world", num)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
