package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutine 起動時にはループが回りきって全部のタスクが再度のタスクとなる
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
