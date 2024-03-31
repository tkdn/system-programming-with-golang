package ch3ioreader

import (
	"archive/zip"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Copy(src string, dist string) {
	file, err := os.Open(src)
	if err != nil {
		fmt.Printf("can't open file: %v", err)
	}
	newFile, err := os.Create(dist)
	if err != nil {
		fmt.Printf("can't create file: %v", err)
	}

	defer func() {
		if err := newFile.Close(); err != nil {
			fmt.Printf("can't close file: %v", err)
		}
	}()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("can't close file: %v", err)
		}
	}()

	io.Copy(newFile, file)
}

func PadRand1024(dist string) int64 {
	file, err := os.Create(dist)
	if err != nil {
		fmt.Printf("can't create file: %v", err)
		return -1
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("can't close file: %v", err)
		}
	}()
	wsize, err := io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		fmt.Printf("can't CopyN: %v", err)
		return -1
	}
	return wsize
}

func CreateZip(dist string, files ...string) {
	zipFile, err := os.Create(dist)
	if err != nil {
		fmt.Printf("can't create file: %v", err)
	}
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for i, f := range files {
		w, err := zipWriter.Create(f)
		if err != nil {
			fmt.Printf("zipWriter can't create: %v", err)
		}
		io.Copy(w, strings.NewReader(fmt.Sprintf("Zipファイルに含まれる内容その%v", i)))
	}
}

func ZipServe() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=test.zip")

		files := []string{"a.txt", "b.txt"}
		zw := zip.NewWriter(w)
		defer zw.Close()

		for i, f := range files {
			w, err := zw.Create(f)
			if err != nil {
				fmt.Printf("zipWriter can't create: %v", err)
				status = http.StatusInternalServerError
			}
			io.Copy(w, strings.NewReader(fmt.Sprintf("Zipファイルに含まれる内容その%v", i)))
		}
		w.WriteHeader(status)
	})

	http.Handle("/zip", handler)
	http.ListenAndServe(":8080", nil)
}
