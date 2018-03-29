package main

import (
	"fmt"
	"math/rand"
)

type ticket struct {
	name  string
	from  string
	to    string
	time  string
	price float64
}

var sem = make(chan ticket, 1)
var done = make(chan int)
var total int = 100

func bookTickets(n int, t ticket) {
	sem <- t
	if n <= total {
		total -= n
		fmt.Printf("i booked %d tickets,there are %d tickets left\n", n, total)
		if total == 0 {
			done <- 1
		}
	}
	<-sem
}

func main() {
	sli := []int{1, 2, 3, 4, 5}
	t := ticket{"wang", "beijing", "shanghai", "today", 1800}
	go func() {
		for total >= 0 {
			go bookTickets(sli[rand.Intn(5)], t)
		}
	}()

	<-done
}
