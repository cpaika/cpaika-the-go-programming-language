package main

import "fmt"

func main() {

  //weird weird weird.  .0 marks them as float64 automatically?
  const freezingF, boilingF = 32.0, 212.0
  fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
  fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
}

// Dumb name why not call it farenheightToCelsius?
//A return statement in the wild!
func fToC(f float64) float64 {
  return (f-32) * 5/9
}

