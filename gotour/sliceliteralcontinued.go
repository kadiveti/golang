package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []string{"ab", "bc", "cd", "de", "ef"}

	c := []struct {
		i int
		q string
	}{
		{1, "ab"},
		{2, "bc"},
		{3, "cd"},
		{4, "de"},
		{5, "ef"},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
