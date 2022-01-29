package main

import (
	"fmt"
	"math"
)

func main() {
	var d, v int

	var tests []int
	fmt.Scanf("%d", &d)
	for i := 0; i < d; i++ {
		fmt.Scanf("%d", &v)
		tests = append(tests, v)
	}
	for _, j := range tests {
		fmt.Println(primeTest(j))
	}

}

func primeTest(p int) string {

	if p == 1 {
		return "NIE"
	}
	if p == 2 {
		return "TAK"
	}
	if p%2 == 0 {
		return "NIE"
	}

	sqrt := math.Sqrt(float64(p))
	b := int(sqrt)

	for i := 3; i <= b; i = i + 2 {
		if p%i == 0 {
			return "NIE"
		}
	}
	return "TAK"
}
