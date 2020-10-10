package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z -= (math.Pow(z, 2) - x) / (math.Pow(2, z))
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
