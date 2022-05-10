package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	t = time.Now()
	fmt.Printf("#%v", t)
}
