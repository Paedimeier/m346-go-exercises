package main

import "fmt"


func computeGrade(points, maxpoints float64) float64 {
	return (points / maxpoints) * 5 + 1
}

func main() {
	fmt.Println(computeGrade(10, 10)) 
	fmt.Println(computeGrade(19, 30))
}
