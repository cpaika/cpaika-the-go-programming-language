package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!!!!", "I can't believe you would run this!")
	var firstArgument = os.Args[1]
	fmt.Println("Hey", firstArgument)
}
