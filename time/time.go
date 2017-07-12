package main

import (
	"time"
	"fmt"
	"log"
)

func doit() {
	fmt.Println("do stuff")
}


// main
func main()  {
	// const layout = "2006.01.02 15:04"
	// const start = "2016.11.04 07:50"

	const layout = "15:04"
	const start = "15:09"
	const interval = "2s"

	loc, _ := time.LoadLocation("Europe/Vienna")
	t, err := time.ParseInLocation(layout, start, loc); if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	today := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc )


	if today.Before(now) {
		fmt.Printf("Schedule %v is in past, moving to next day\n", today)
		today = today.Add(time.Hour*24)
	}

	fmt.Printf("Waiting until %v\n", today)

	waittime := today.Sub(now)

	fmt.Printf("Waiting %v to start\n", waittime)

	time.Sleep(waittime)

	duration,_ := time.ParseDuration(interval)
	ticker := time.NewTicker(duration)

	for {
		time := <- ticker.C
		fmt.Printf("Running at %v\n", time)
		doit()
	}


	// D := t.Sub(time.Now())
	// fmt.Println(d)

	// time.AfterFunc(d, func() {
	// 	fmt.Println("expired")
	// })

	// dc := make(chan time.Duration)
	// time.AfterFunc(d, schedule)


	// fmt.Println(t.Sub(time.Now()))

	// time.Sleep(time.Hour)
}
