package main

import "fmt"

func add(x, y int) int { // same as `func add(x int, y int) int`
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
