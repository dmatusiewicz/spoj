package main

import (
	"fmt"
	"math"
)

func main() {

	var r, d float64
	fmt.Scanf("%f %f", &r, &d)
	fmt.Printf("%.2f", poleKola(r, d))
}

func poleKola(r, d float64) float64 {
	powR := math.Pow(r, 2)
	powD := math.Pow(d, 2)
	return math.Pi * (powR - powD/4)
}
