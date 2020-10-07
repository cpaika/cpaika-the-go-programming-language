//echo4 prints its command line args
package main

import (
  "flag"
  "fmt"
  "strings"
)

//I guess flag is a built in object/function
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "", "seperator")

func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }
}
