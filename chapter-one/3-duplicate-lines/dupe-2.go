package main

//parenthesis for import
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	//Make a slice with the arguments after the initial program command
	files := os.Args[1:]
	//go doesnt have parenthesis around if statements
	if len(files) == 0 {
		// the real question - is counts passed by reference?  It must be since its an object
		countLines(os.Stdin, counts)
	} else {
		//now we have to load the file from disk
		// we dont care about i so we drop it with the blank identifier_
		// why isnt range a function?
		for _, arg := range files {
			//Open the file
			// Why oh why does golang not have a dedicated error channel?

			f, err := os.Open(arg)
			//nil must be the equivalent of null obviously
			if err != nil {
				// Print to stderr if there is an error
				// oh fancy fprintf I wonder what that is
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// what is continue thats the real question
				continue
			}

			countLines(f, counts)
			//Always close your file handlers!
			f.Close()
		}
	}
	printDupes(counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printDupes(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d  %s\n", n, line)
		}
	}
}

//golang has a built in error type
//functions may be declared in any order in golang
// a map is a refernce to the data structure created by make
//Passing it to a function, the function gets a copy of the reference so any action modifies the underlying data structure
//We are using streaming mode here, which is better than reading everything in one big gulp
