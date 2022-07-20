package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func min(x int, y int) {
	if x < y {
		return x
	} else {
		return y
	}
}
