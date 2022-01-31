package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var d int
	var data []string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d", &d)

	for i := 0; i < d; i++ {
		scanner.Scan()
		data = append(data, scanner.Text())

	}

	for _, j := range data {
		revertData(j)
	}
}

func revertData(d string) {
	splitted := strings.SplitN(d, " ", -1)

	for i := len(splitted) - 1; i > 0; i-- {
		fmt.Printf("%s ", splitted[i])
	}
	fmt.Printf("\n")

}
