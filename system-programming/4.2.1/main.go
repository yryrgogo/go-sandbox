package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		done <- true
	}()
	// 上の goroutine を削除すると、deadlock になり compile error になる
	<-done
	fmt.Println("all tasks are finished")
}
