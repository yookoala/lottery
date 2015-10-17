package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// total number of available person in the lottery
var total = flag.Int("total", 1000, "Total number of people in the lottery")

func init() {
	flag.Parse()
}

func uniqueInt(in <-chan int) <-chan int {
	history := make(map[int]bool)
	out := make(chan int)
	go func() {
		for {
			n := <-in
			_, ok := history[n]
			for ; ok; _, ok = history[n] {
				n = <-in
			}
			history[n] = true
			out <- n
		}
	}()
	return out
}

func randNumbers(n int) <-chan int {
	rand.Seed(time.Now().Unix())
	out := make(chan int)
	go func() {
		for {
			out <- rand.Intn(n) + 1
		}
	}()
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for n := range uniqueInt(randNumbers(*total)) {
		fmt.Printf("%d", n)
		reader.ReadString('\n')
	}
}
