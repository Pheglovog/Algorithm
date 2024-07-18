package main

import (
	"fmt"
	"strings"
	"time"
)

func timeTest() {
	tryAfter()
	time.Sleep(1 * time.Second)
	tryDuration()
	tryLocation()
}

func tryAfter() {
	fmt.Println(strings.Repeat("-", 8) + "tryAfter" + strings.Repeat("-", 8))
	handle := func(int) {}
	c := make(chan int)
	defer close(c)
	select {
	case m := <-c:
		handle(m)
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	}
}

func tryDuration() {
	fmt.Println(strings.Repeat("-", 8) + "tryDuration" + strings.Repeat("-", 8))
	expensiveCall := func() {}
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("t0: %v\n", t0)
	fmt.Printf("t1: %v\n", t1)
	fmt.Printf("t1.Sub(t0): %v\n", t1.Sub(t0))
}

func tryLocation() {
	fmt.Println(strings.Repeat("-", 8) + "tryLocation" + strings.Repeat("-", 8))
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")

	// Creating a time requires a location. Common locations are time.Local and time.UTC.
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)
	fmt.Printf("sameTimeInBeijing.In(time.UTC): %v\n", sameTimeInBeijing.In(time.UTC))
}
