package main

import (
	"bufio"
	"fmt"
	"os"
)

//ctrl-D is the end of input signal in posix terminals

func main() {
	//make is a built in function, it will be discussed later.  For now I'm a little confused/curious
	counts := make(map[string]int)
	//short variable declaration
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// Iterating over a map is random, and the order will be different every time
	for line, n := range counts {
		if n > 1 {
			// Printf verbs
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
