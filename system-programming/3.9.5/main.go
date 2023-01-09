package main

import (
	"bytes"
	"io"
	"os"
)

func CopyN1(dest io.Writer, src io.Reader, length int) {
	b := make([]byte, length)
	src.Read(b)
	dest.Write(b)
}

// これは失敗する
func CopyN2(dest io.Writer, src io.Reader, length int) {
	b := make([]byte, length)
	buf := bytes.NewBuffer(b)
	io.Copy(buf, src)
	io.Copy(dest, buf)
}

func main() {
	b := bytes.NewBuffer(
		[]byte{
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
			0x40,
		},
	)
	CopyN1(os.Stdout, b, 3)
	// CopyN2(os.Stdout, b, 3)
}
