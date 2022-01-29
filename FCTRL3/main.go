package main

import (
	"fmt"
)

func main() {
	var tests []uint
	var d, i uint8
	var tc uint
	fmt.Scanf("%d", &d)
	for i = 0; i < d; i++ {
		fmt.Scanf("%d", &tc)
		tests = append(tests, tc)
	}

	for _, f := range tests {
		if f > 9 {
			fmt.Println("0 0")
		} else {
			digitPrint(factor(f))
		}

	}

}

func digitPrint(f uint) {
	fmt.Printf("%d %d\n", (f/10)%10, f%10)
}

func factor(f uint) uint {
	var r uint = 1
	var i uint
	for i = 1; i <= f; i++ {
		r *= i
	}
	return r
}
