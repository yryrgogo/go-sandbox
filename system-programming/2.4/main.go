package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("Tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

func httpGet() {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	conn, err := net.Dial("Tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	req.Write(conn)
}
