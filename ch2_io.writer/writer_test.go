package ch2iowriter_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	ch2iowriter "github.com/tkdn/system-programming-with-golang/ch2_io.writer"
)

func TestFileWrite(t *testing.T) {
	path := "./.temp/file.txt"
	if err := ch2iowriter.FileWrite(path); err != nil {
		t.Errorf("cant' FileWrite: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			t.Errorf("can't remove file %v: %v", path, err)
		}
	})
}

func TestWriteCSV(t *testing.T) {
	var buff bytes.Buffer
	records := [][]string{
		{"id", "name", "age"},
		{"123", "bob", "22"},
	}
	ch2iowriter.WriteCSV(&buff, records)
	want := "id,name,age\n123,bob,22\n"
	got := buff.String()
	fmt.Println(want)
	fmt.Println(got)
	if want != got {
		t.Fatalf("not matched:\n  want:\n%v\n  got:\n%v", want, got)
	}
}
