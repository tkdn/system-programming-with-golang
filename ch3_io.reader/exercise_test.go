package ch3ioreader_test

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	ch3ioreader "github.com/tkdn/system-programming-with-golang/ch3_io.reader"
)

func TestCopy(t *testing.T) {
	want := "FooBar2000\nSample text."
	old := "./.temp/old.txt"
	new := "./.temp/new.txt"
	createTextFile(old, want)
	ch3ioreader.Copy(old, new)
	f, err := os.Open(new)
	if err != nil {
		t.Error(err)
	}
	got, err := io.ReadAll(f)
	if err != nil && string(got) != want {
		t.Error(err)
	}
}

func TestPadRand(t *testing.T) {
	got := ch3ioreader.PadRand1024("./.temp/rand")
	if got != 1024 {
		t.Errorf("not match size: %#v", got)
	}
}

func createTextFile(dist string, text string) {
	file, err := os.Create(dist)
	if err != nil {
		fmt.Printf("can't create file: %v", err)
	}
	fmt.Fprint(file, text)
	if err := file.Close(); err != nil {
		fmt.Printf("can't close file: %v", err)
	}
}

func TestCreateZip(t *testing.T) {
	files := []string{"a.txt", "b.txt"}
	zipFile := "./.temp/test.zip"
	ch3ioreader.CreateZip(zipFile, files...)

	zr, _ := zip.OpenReader(zipFile)
	defer zr.Close()

	for i, f := range zr.File {
		if f.Name != files[i] {
			t.Errorf("file name not mathed: want: %#v, but got: %#v", files[i], f.Name)
		}
	}
}

func TestCopyN(t *testing.T) {
	w := os.Stdout
	size, err := ch3ioreader.CopyN(w, rand.Reader, 64)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if size != 64 {
		t.Errorf("size is not matched: %v", err)
	}
}

func TestStream(t *testing.T) {
	var buff bytes.Buffer
	want := "ASCII"
	stream := ch3ioreader.Stream()
	buff.ReadFrom(stream)
	got := buff.String()
	if want != got {
		t.Errorf("not matched... \nwant: %v\nbut got: %v", want, got)
	}
}

func TestMain(m *testing.M) {
	files, _ := filepath.Glob("./.temp/*")
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			fmt.Printf("can't rm file: %v", err)
		}
	}
	status := m.Run()
	os.Exit(status)
}
