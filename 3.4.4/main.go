package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30, 0x40})
	buffer3 := bytes.NewBufferString("初期文字列")
	fmt.Printf("%v\n", buffer1)
	fmt.Printf("%v\n", buffer2)
	fmt.Printf("%v\n", buffer3)
}
