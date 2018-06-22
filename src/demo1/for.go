package main

import "fmt"

func loop1() {
	for i := 3; i <= 7; i++ {
		fmt.Println(i)
		// return i
	}
}

func loop2() {
	for i := 3; i <= 7; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}
}

func loop3() {
	for i := 3; i <= 7; i++ {
		if i == 4 {
			continue
		}
		fmt.Println(i)
	}
}
