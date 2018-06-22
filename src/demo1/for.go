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

func loop4() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func loop5() {
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
