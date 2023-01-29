package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/yryrgogo/go-sandbox/jsonrpc/jsonrpc2"
)

func handle(
	ctx context.Context,
	conn *jsonrpc2.Conn,
	req *jsonrpc2.Request,
) (result interface{}, err error) {
	fmt.Println("jsonrpc2 handler was called!")

	return nil, &jsonrpc2.Error{
		Code:    jsonrpc2.CodeMethodNotFound,
		Message: fmt.Sprintf("method not supported: %s", req.Method),
	}
}

func start() {
	handler := jsonrpc2.HandlerWithError(handle)
	connOpt := []jsonrpc2.ConnOpt{}

	fmt.Println("start!")
	<-jsonrpc2.NewConn(
		context.Background(),
		jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		handler,
		connOpt...,
	).DisconnectNotify()
}

type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error) {
	size, err := os.Stdin.Read(p)
	fmt.Printf("stdrwc.Read input: %v\n", string(p))
	return size, err
}

func (stdrwc) Write(p []byte) (int, error) {
	fmt.Printf("write %v\n", string(p))
	return os.Stdout.Write(p)
}

func (stdrwc) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	return os.Stdout.Close()
}

func play_reader(conn io.ReadWriteCloser) {
	reader := bufio.NewReader(conn)

	fmt.Println("play reader start!")
	line, err := reader.ReadString('\r')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("line: %v\n", line)
}

// func play_writer(conn io.ReadWriteCloser) {
// 	writer := bufio.NewWriter(conn)
//
// 	writer
// }

func main() {
	// start()
	play_reader(stdrwc{})
}
