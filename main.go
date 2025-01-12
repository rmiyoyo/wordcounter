package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	words := flag.Bool("w", false, "Count words")
	chars := flag.Bool("c", false, "Count characters")
	flag.Parse()

	if !*lines && !*words && !*chars {
		*lines = true
	}

	fmt.Println(count(os.Stdin, *lines, *words, *chars))
}

func count(r io.Reader, countLines, countWords, countChars bool) int {
	scanner := bufio.NewScanner(r)

	if !countWords {
		scanner.Split(bufio.ScanWords)
	} else if !countLines {
		scanner.Split(bufio.ScanRunes)
	}

	wc := 0

	for scanner.Scan() {
		if countChars {
			wc += len(strings.TrimSpace(scanner.Text()))
		} else {
			wc++
		}
	}

	return wc
}
