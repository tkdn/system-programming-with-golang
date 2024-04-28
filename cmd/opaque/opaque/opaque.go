package opaque

import "fmt"

type Opaque interface {
	GetNumber() int
	implementsOpaque()
}

func DoSomethingWithOpaque(o Opaque) string {
	return fmt.Sprintf("Hello opaque #%d", o.GetNumber())
}
