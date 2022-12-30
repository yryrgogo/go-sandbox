package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	// i := make([]byte, 4)
	var i int32
	binary.Read(bytes.NewBuffer(data), binary.BigEndian, &i)
	fmt.Printf("data: %d %d\n", i, data)
}
