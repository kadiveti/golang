package main

import (
	"fmt"
)

type car struct {
	gas   int
	speed int
	cost  int
	color string
}

func main() {
	a := car{1, 2, 3, "blue"}
	fmt.Println(a.color)
}
