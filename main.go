package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// total number of available person in the lottery
var total = flag.Int("total", 0, "Total number of people in the lottery")

func init() {
	flag.Parse()
	if *total <= 0 {
		fmt.Print("Usage: lottery -total [NUMBER]\n\n")
		os.Exit(1)
	}
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
	url := "https://github.com/yookoala/lottery"
	fmt.Printf("\nSource code: %s\n\nTotal number: %d\n"+
		"=================\n\n",
		url, *total)

	for n := range uniqueInt(randNumbers(*total)) {
		t := time.Now().Format("15:04:05.999")
		t += strings.Repeat("0", 12-len(t))
		fmt.Printf("[%s] %d", t, n)
		reader.ReadString('\n')
	}
}
