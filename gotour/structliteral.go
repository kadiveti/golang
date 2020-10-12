package main

import (
	"fmt"
)

func main() {

	names := [3]string{
		"Rahul",
		"Don",
		"JOhn",
	}

	fmt.Println(names)

	a := names[1:2]
	b := names[0:]

	fmt.Println(a, b)

	b[0] = "xxxxx"
	fmt.Println(a, b)
	fmt.Println(names)
}
