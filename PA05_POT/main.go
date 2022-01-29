package main

import (
	"fmt"
	"math"
	"strconv"
)

type testCase struct {
	a, b float64
}

func main() {
	var d, i uint8
	// var a, b float64
	var testCases []testCase

	fmt.Scanf("%d", &d)
	for i = 0; i < d; i++ {
		var ttemp testCase

		fmt.Scanf("%f %f", &ttemp.a, &ttemp.b)
		testCases = append(testCases, ttemp)
		// fmt.Printf("%s %1.f %1.f\n", power(a, b), a, b)
	}

	for _, j := range testCases {
		fmt.Println(power(j.a, j.b))
	}
}

func power(a, b float64) string {

	c := math.Mod(a, 10)
	if c == 0.0 {
		return string("0")
	}
	if c == 1.0 {
		return string("1")
	}
	var res float64
	x := math.Mod(b, 4.0)
	if x == 0 {
		res = math.Pow(c, 4)
	} else {
		res = math.Pow(c, x)
	}

	if res > 9.0 {
		s := strconv.Itoa(int(res))
		l := len(s)

		return s[l-1:]
	}
	return strconv.Itoa(int(res))
}
