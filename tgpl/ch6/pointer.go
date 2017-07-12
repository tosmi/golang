package main

import (
	"fmt"
)

// Point is a point
type Point struct{ x, y float64 }

func main() {
	p := Point{1, 2}
	fmt.Printf("Point: %v", p)

}
