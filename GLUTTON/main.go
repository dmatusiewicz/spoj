package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

type party struct {
	cookieBoxSize int64
	participants  []participant
}

type participant struct {
	cookieEatingTime int64
}

func main() {
	var d int
	var parties []party

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	d, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < d; i++ {
		scanner.Scan()
		p := getPartyDetails(scanner.Text())
		parties = append(parties, p)
	}

	for _, j := range parties {
		fmt.Printf("%d\n", countBoxes(j))
	}

}

func countBoxes(p party) int64 {
	var boxes, cookieCount int64
	const day int64 = 86400
	for _, j := range p.participants {
		cookieCount += day / j.cookieEatingTime
	}
	if cookieCount%p.cookieBoxSize != 0 {
		boxes++
	}
	boxes += cookieCount / p.cookieBoxSize

	return boxes
}

func getPartyDetails(s string) party {
	pd := strings.Split(s, " ")
	cookieBoxSize, err := strconv.ParseInt(pd[1], 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	p := party{
		cookieBoxSize: cookieBoxSize,
	}

	people, err := strconv.ParseInt(pd[0], 0, 32)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < int(people); i++ {
		scanner.Scan()
		cookieEatingTime, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		participant := participant{
			cookieEatingTime: cookieEatingTime,
		}
		p.participants = append(p.participants, participant)

	}

	return p
}
