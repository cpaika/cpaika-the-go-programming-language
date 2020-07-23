package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
	//fmt.Println(output)
}
