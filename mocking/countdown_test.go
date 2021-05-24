package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type CountdownOperationsSpy struct {
	calls []string
}

func (cs *CountdownOperationsSpy) Sleep() {
	cs.calls = append(cs.calls, sleep)
}

func (cs *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	cs.calls = append(cs.calls, write)
	return
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("sleep before print", func(t *testing.T) {
		cs := &CountdownOperationsSpy{}

		Countdown(cs, cs)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, cs.calls) {
			t.Errorf("wanted calls %v got %v", want, cs.calls)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &spyTime{}
	sleeper := &configurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("want: %v, got: %v", sleepTime, spyTime.durationSlept)
	}
}
