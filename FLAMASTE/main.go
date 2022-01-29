package main

import "fmt"

func main() {

	var d int
	var v string

	var tests []string
	fmt.Scanf("%d", &d)
	for i := 0; i < d; i++ {
		fmt.Scanf("%s", &v)
		tests = append(tests, v)
	}
	for _, j := range tests {
		convert(j)
	}
}

func convert(s string) {
	l := len(s)
	// var shrtS string
	var counter int = 1
	for i := 0; i < l; i++ {

		// Jak obecny jest taki sam jak nastepny s = ZX s[0] = Z s[1] = X
		if i+1 < l && s[i] == s[i+1] {
			counter++
		}

		// Jak obecny jest inny od nastepnego a licznik jest jeden
		if i+1 < l && s[i] != s[i+1] && counter == 1 {
			fmt.Printf("%c", s[i])
		}

		// Jak obecny jest inny od nastepnego a licznik nie jest wiekszy od jeden i zresetuj
		if i+1 < l && s[i] != s[i+1] && counter > 1 {
			if counter == 2 {
				fmt.Printf("%c%c", s[i-1], s[i])
			} else {
				fmt.Printf("%c%d", s[i], counter)
			}
			counter = 1
		}

		// Jak dotarlismy do konca i trzeba wypisac co jest
		if i+1 == l {
			if counter > 1 {
				if counter == 2 {
					fmt.Printf("%c%c", s[i-1], s[i])
				} else {
					fmt.Printf("%c%d", s[i], counter)
				}
			}
			if counter == 1 {
				fmt.Printf("%c", s[i])
			}
			fmt.Printf("\n")
			counter = 1

		}

	}
}
