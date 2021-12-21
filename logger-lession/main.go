package main

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")
	// 记完日志后会调用os.Exit(1)
	// log.Fatal("Bye.")
	// 记完日志后会调用 panic()
	log.Panic("I'm bailing.")
}
