package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

func computeHypotenuse(kathete1, kathete2 float64) float64 {
	return math.Sqrt(math.Pow(kathete1, 2) + math.Pow(kathete2, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println(computeHypotenuse(3, 4))
	fmt.Println(computeHypotenuse(5, 12))
}
