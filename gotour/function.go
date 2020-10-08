package main
// int can be passed in two ways

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func main() {
	fmt.Println(add(2, 3))
	fmt.Println(sub(5, 3))
}
