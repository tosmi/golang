package main

import (
	"time"
	"fmt"
	"log"
)

func doit() {

	fmt.Println("do stuff")



}

func schedule() {
	fmt.Println("in schedule")

	doit()


}



// main
func main()  {
	const layout = "2006.01.02 15:04"
	const start = "2016.11.04 07:50"
	const interval = "2s"

	loc, _ := time.LoadLocation("Europe/Vienna")

	t,err := time.ParseInLocation(layout, start, loc); if err != nil {
		log.Fatal(err)
	}

	duration := time.ParseDuration(interval)

	if t.Before(time.Now()) {
		fmt.Println("Schedule start is before now!")
	}

	d := t.Sub(time.Now())
	fmt.Println(d)

	time.AfterFunc(d, func() {
		fmt.Println("expired")
	})

	dc := make(chan time.Duration)
	time.AfterFunc(d, schedule)


	fmt.Println(t)
	fmt.Println(t.Sub(time.Now()))

	time.Sleep(time.Hour)
}
