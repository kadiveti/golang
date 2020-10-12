package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	printslice(s)

	s = s[:0]
	printslice(s)

	s = s[:4]
	printslice(s)

	s = s[:2]
	printslice(s)

}

func printslice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
