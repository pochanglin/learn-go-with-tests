package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMsg := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Allen", "")
		want := "Hello, Allen"
		assertCorrectMsg(t, got, want)
	})

	t.Run("empty string default to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Allen", "Spanish")
		want := "Hola, Allen"
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Allen", "French")
		want := "Bonjour, Allen"
		assertCorrectMsg(t, got, want)
	})

}
