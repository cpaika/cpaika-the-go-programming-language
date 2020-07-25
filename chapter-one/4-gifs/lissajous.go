package main

import (
	"image/color"
)

// building the variable palette, which is an array of two colors
var palette = []color.Color{color.White, color.Black}

//first and second index of the palette
// I dont love how this works
// insteresting way to declare constants tbh I don't really get it
const (
	whiteIndex = 0
	blackIndex = 1
)
