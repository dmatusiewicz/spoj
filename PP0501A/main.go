package main

import "fmt"

type liczby struct {
	a, b int
}

func main() {
	var d int
	var v liczby

	var tests []liczby

	fmt.Scanf("%d", &d)
	for i := 0; i < d; i++ {
		fmt.Scanf("%d %d", &v.a, &v.b)
		tests = append(tests, v)
	}
	for _, j := range tests {
		fmt.Printf("%d\n", nwd(j.a, j.b))
	}
}

func nwd(a, b int) int {
	if b == 0 {
		return a
	}
	return nwd(b, a%b)
}
