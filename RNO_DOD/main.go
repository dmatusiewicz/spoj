package main

import "fmt"

type adder struct {
	n int
	a []int
}

func main() {
	var tC int
	var atC []adder
	var a adder

	fmt.Scanf("%d", &tC)
	for i := 0; i < tC; i++ {
		a = adder{}
		fmt.Scanf("%d", &a.n)
		for ii := 0; ii < a.n; ii++ {
			var z int
			fmt.Scanf("%d", &z)
			a.a = append(a.a, z)
		}
		atC = append(atC, a)
	}

	for _, j := range atC {
		fmt.Printf("%d\n", addAll(j))
		// addAll(j)
	}
}

func addAll(a adder) int {
	var sum int
	for _, j := range a.a {
		sum += j
	}
	return sum
}
