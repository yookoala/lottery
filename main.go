package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/yookoala/lottery/reader"
)

// min and max random number of the lottery
var max = flag.Int("max", 0, "Max random number in the lottery")
var min = flag.Int("min", -1, "Min random number in the lottery")
var filename = flag.String("file", "", "Filename of the lottery information")

func init() {
	flag.Parse()
	if *max <= 0 && *filename == "" {
		fmt.Print("Usage: lottery -max <NUMBER> [-min <NUMBER>]\n" +
			"or:    lottery -file <.XLSX FILENAME>\n\n")
		os.Exit(1)
	}

	if *filename != "" {
		// file mode defaults
		if *min < 0 {
			*min = 2 // default skip header
		}

	} else {
		if *min <= 0 {
			*min = 1
		}
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

func randInts(min, max int) <-chan int {
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

func padNumToStr(pad int) func(int) (string, error) {
	fmtStr := fmt.Sprintf("%%%dd", pad)
	return func(n int) (str string, err error) {
		str = fmt.Sprintf(fmtStr, n)
		return
	}
}

func formatNum(fn func(int) (string, error), in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for {
			n := <-in
			str, err := fn(n)
			if err != nil {
				panic(err)
			}
			out <- str
		}
	}()
	return out
}

func main() {

	var seq <-chan string

	if *filename != "" {

		// file mode
		c, err := reader.OpenXLSXSheet(*filename, 0)
		if err != nil {
			panic(err)
			return
		}

		if *max == 0 {
			*max = c.Len() // last row
		}

		fmter := func(n int) (str string, err error) {
			strs, err := reader.ReadMulti(c, 0, 1)(n)
			if err != nil {
				return
			}
			str = strings.Join(strs, ", ")
			return
		}
		seq = formatNum(fmter, uniqueInt(randInts(*min-1, *max-1)))
	} else {
		// raw number mode
		l := len(fmt.Sprintf("%d", *max))
		seq = formatNum(padNumToStr(l), uniqueInt(randInts(*min, *max)))
	}

	url := "https://github.com/yookoala/lottery"
	fmt.Printf("\nSource code: %s\n\n", url)
	fmt.Printf("Max: %d\n", *max)
	fmt.Printf("Min: %d\n", *min)
	fmt.Printf("=================\n\n")

	reader := bufio.NewReader(os.Stdin)
	for n := range seq {
		t := time.Now().Format("15:04:05.999")
		t += strings.Repeat("0", 12-len(t))
		fmt.Printf("[%s] %s", t, n)
		reader.ReadString('\n')
	}
}
