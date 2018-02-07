// Prints the count and text of each line that appears
// more than once in the input. It reads from stdin or
// from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open returns 2 values,
			// the opened file and an error
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// Close closes the file when the last line is read
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Map is a reference to the data structure created by make
// When a map is passed to a func, the func receives that references,
// so any changes to underlying data that this func makes to map
// are also visible to other funcs right away
// Map here is a 'reference type' in Swift terms, different from NSDictionary
// or Dictionary
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Ignoring errors from input.Err() for the sake of brevity
}
