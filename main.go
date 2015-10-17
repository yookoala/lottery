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

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())
	for {
		var result = rand.Intn(*total) + 1
		fmt.Printf("%d", result)
		reader.ReadString('\n')
	}
}
