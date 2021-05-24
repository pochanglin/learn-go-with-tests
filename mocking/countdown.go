package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type defaultSleeper struct{}

func (d *defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type configurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *configurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, s Sleeper) {
	const finalWord = "Go!"
	const countdownStart = 3
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &configurableSleeper{
		duration: 2 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
