package main

import (
	"fmt"
)

type value struct {
	x int
	y int
}

func main() {
	v := value{1, 2}
	p := &v
	p.x = 1e9
	fmt.Println(*p)
}
