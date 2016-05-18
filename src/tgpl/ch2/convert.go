package main

import (
	"tgpl/ch2/tempconv"
	"tgpl/ch2/lengthconv"
	"tgpl/ch2/weightconv"
	"fmt"
)

func main() {
	k := weightconv.Kilogram(1)
	m := lengthconv.Meter(1)
	c := tempconv.Celcius(1)

	fmt.Printf("%g %g %g", k, m, c)
}
