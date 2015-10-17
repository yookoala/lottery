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

// min and max random number of the lottery
var max = flag.Int("max", 0, "Max random number in the lottery")
var min = flag.Int("min", 1, "Min random number in the lottery")

func init() {
	flag.Parse()
	if *max <= 0 {
		fmt.Print("Usage: lottery -max <NUMBER> [-min <NUMBER>]\n\n")
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

func randNumbers(min, max int) <-chan int {
	rand.Seed(time.Now().Unix())
	out := make(chan int)
	n := max - min + 1
	go func() {
		for {
			out <- rand.Intn(n) + min
		}
	}()
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	url := "https://github.com/yookoala/lottery"
	fmt.Printf("\nSource code: %s\n\n", url)
	fmt.Printf("Max: %d\n", *max)
	if *min != 1 {
		fmt.Printf("Min: %d\n", *min)
	}
	fmt.Printf("=================\n\n")

	for n := range uniqueInt(randNumbers(*min, *max)) {
		t := time.Now().Format("15:04:05.999")
		t += strings.Repeat("0", 12-len(t))
		fmt.Printf("[%s] %d", t, n)
		reader.ReadString('\n')
	}
}
