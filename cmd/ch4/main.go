package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Second * 5
	timeout := time.After(s)
	<-timeout
	fmt.Println("after 5 secnods")
}
