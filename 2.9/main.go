package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Helo": "World",
	}

	multiW := io.MultiWriter(w, os.Stdout)

	writer := gzip.NewWriter(multiW)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
	writer.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
