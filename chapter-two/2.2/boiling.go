package main
import "fmt"
//package level since lowercase
const boilingF = 212.0

func main() {
  var f = boilingF
  var c = (f - 32) * 5/9
  fmt.Printf("boiling point %gF or %gC\n", f, c)
}
