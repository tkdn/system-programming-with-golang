package ch3ioreader

import (
	"fmt"
	"strings"
)

var d = "123 1.234 1.0e4 test"

func ReadData() {
	r := strings.NewReader(d)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(r, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
