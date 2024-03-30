package ch2iowriter

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "json/application")
	gzp := gzip.NewWriter(w)
	mwriter := io.MultiWriter(gzp, os.Stdout)
	source := map[string]string{
		"Hello": "World",
	}
	json.NewEncoder(mwriter).Encode(source)
	gzp.Flush()
}

func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
