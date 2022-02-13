package main

import "fmt"

type dataManipulator struct {
	a, b string
}

func main() {
	var d int
	var data []dataManipulator
	var dataHolder dataManipulator
	fmt.Scanf("%d", &d)
	for i := 0; i < d; i++ {
		fmt.Scanf("%s %s", &dataHolder.a, &dataHolder.b)
		data = append(data, dataHolder)
	}

	for _, j := range data {
		fmt.Printf("%s\n", stringMerge(j))
	}
}

func stringMerge(d dataManipulator) string {
	lana := len(d.a)
	lenb := len(d.b)
	var l int
	var s []rune
	if lana < lenb {
		l = lana
	} else {
		l = lenb
	}

	for i := 0; i < l; i++ {

		s = append(s, rune(d.a[i]))
		s = append(s, rune(d.b[i]))

	}
	merged := string(s)
	return merged
}
