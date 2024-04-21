package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Join
	fmt.Printf("Temp file Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
	// Split
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s,\n", dir, name)
	// Split each path
	fragments := strings.Split("/foo/bar/2000/temp.txt", string(filepath.Separator))
	fmt.Printf("paths: %s\n", fragments)
}
