package main

import (
	"fmt"
	"tgpl/ch2/tempconv"
)

func main() {
	var test tempconv.Celcius = -3
	fmt.Println(test)
	fmt.Println(tempconv.CtoF(test))
	fmt.Println(tempconv.CtoK(test))
}
