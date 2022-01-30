package main

import "fmt"

type klasy struct {
	a, b int
}

func main() {
	var d int
	var v klasy

	var tests []klasy

	fmt.Scanf("%d", &d)
	for i := 0; i < d; i++ {
		fmt.Scanf("%d %d", &v.a, &v.b)
		tests = append(tests, v)
	}
	for _, j := range tests {
		fmt.Printf("%d\n", nww(j))
	}
}

func nww(k klasy) int {
	return (k.a * k.b) / nwd(k.a, k.b)
}

func nwd(a, b int) int {
	if b == 0 {
		return a
	}
	return nwd(b, a%b)
}
