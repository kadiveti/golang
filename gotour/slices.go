package main

import (
	"fmt"
)

func main() {
	p := [5]int{1, 2, 3, 4, 5}

	var a []int = p[1:4]
	fmt.Println(a)
}
