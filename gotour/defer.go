package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("master")

	fmt.Println("Hello")
}
