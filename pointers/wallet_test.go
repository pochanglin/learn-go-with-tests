package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t,w,Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{
			balance: Bitcoin(10),
		}
		err := w.Withdraw(Bitcoin(10))
		assertBalance(t,w,Bitcoin(0))
		assertNoError(t,err)
	})

	t.Run("Withdraw insufficient money", func(t *testing.T) {
		startBalance := Bitcoin(20)
		w := Wallet{
			balance: startBalance,
		}
		err := w.Withdraw(Bitcoin(100))
		assertBalance(t,w,startBalance)
		assertError(t,err,ErrInsufficientFunds)
	})
	
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin){
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, err error, want error){
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t testing.TB, got error){
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
