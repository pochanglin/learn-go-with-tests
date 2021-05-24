package racewebsite

import (
	"net/http"
	"time"
)

func Racer(a string, b string) (w string) {
	aD := measureResponseTime(a)

	bD := measureResponseTime(b)

	if aD < bD {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
