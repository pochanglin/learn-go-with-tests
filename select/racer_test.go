package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowServer := makeDelayedServer(20 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		want := fastServer.URL
		got, err := Racer(fastServer.URL, slowServer.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("return error if a server doesn't response with 10s", func(t *testing.T) {
		serverA := makeDelayedServer(20 * time.Millisecond)

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 15*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			rw.WriteHeader(http.StatusOK)
		}),
	)

}
