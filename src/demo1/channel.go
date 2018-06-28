package main

import "fmt"

// func channeldemo() {
// 	ch := make(chan string)

// 	go func(msg string) {
// 		ch <- msg
// 	}("hello")

// 	recv := <-ch
// 	fmt.Println(recv)
// }

func channeldemo() {
	// buffered
	ch := make(chan string, 1)

	func(msg string) {
		ch <- msg
	}("hello")

	recv := <-ch
	fmt.Println(recv)
}
