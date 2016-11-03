package main

import (
	"time"
	"fmt"
	"os"
)

func run() {
	fmt.Println("called run")

}



// main
func main()  {
	const layout = "2006.01.02 15:04"
	const start = "2016.11.03 17:18"
	const interval = "2"

	loc, _ := time.LoadLocation("Europe/Vienna")

	t,err := time.ParseInLocation(layout, start, loc); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if t.Before(time.Now()) {
		fmt.Println("Schedule start is before now!")
	}

	d := t.Sub(time.Now())
	fmt.Println(d)

	time.AfterFunc(d, func() {
		fmt.Println("expired")
	})

	fmt.Println(t)
	fmt.Println(t.Sub(time.Now()))

	time.Sleep(time.Hour)
}
