package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i
	*p = 25
	fmt.Println(i)

	p = &j
	*p = 97
	fmt.Println(j)
}
