package main

import (
	"fmt"
)

type value struct {
	x, y int
}

var (
	v  = value{1, 2}
	v1 = value{x: 1}
	v2 = value{y: 1}
	p  = &value{1, 2}
)

func main() {
	fmt.Println(v)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(*p)
}
