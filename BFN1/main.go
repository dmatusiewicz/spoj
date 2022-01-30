package main

import (
	"fmt"
	"math"
)

func main() {

	var t, v int
	var tests []int
	var tries int

	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &v)
		tests = append(tests, v)
	}

	for _, j := range tests {
		for {
			if palindrome(j) {
				fmt.Printf("%d %d\n", j, tries)
				tries = 0
				break
			} else {
				j += revertNum(j)
				tries++
			}
		}
	}
}

func revertNum(i int) int {
	var r int
	for {
		r = r*10 + i%10
		i /= 10
		if i > 0 {
			continue
		} else {
			return r
		}
	}
}

func countDigits(i int) int {
	counter := 1
	for {
		i /= 10
		if i != 0 {
			counter++
		} else {
			break
		}
	}
	return counter
}

func palindrome(i int) bool {
	l := countDigits(i)
	var x1, x2 int
	for j := 0; j < l/2; j++ {
		if j == 0 {
			x1 = i / int(math.Pow(10.0, float64(l-1-j)))
		} else {
			x1 = (i % int(math.Pow(10.0, float64(l-j)))) / int(math.Pow(10.0, float64(l-1-j)))
		}
		x2 = (i % int(math.Pow(10.0, float64(j+1)))) / int(math.Pow(10.0, float64(j)))
		if x1 != x2 {
			return false
		}
	}
	return true
}
