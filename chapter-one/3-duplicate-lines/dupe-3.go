package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	printDupes(counts)
}

// Need to be the main 3 package since otherwise this public function conflicts with another in the package.  Interesting...
// The package being dupe.go and dupe-2.go, which are both main.
func printDupes(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d  %s\n", n, line)
		}
	}
}
