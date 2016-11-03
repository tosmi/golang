package main

import (
	"fmt"
	"os"
	"strconv"
	"tgpl/ch2/tempconv"
	"tgpl/ch2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celcius(t)
		ke := tempconv.Kelvin(t)
		k := weightconv.Kilogram(t)
		p := weightconv.Pounds(t)

		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
			f, tempconv.FtoC(f),
			c, tempconv.CtoF(c),
			ke, tempconv.KtoC(ke),
			k, weightconv.KtoP(k),
			p, weightconv.PtoK(p),
		)
	}
}
