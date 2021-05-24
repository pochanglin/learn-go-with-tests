package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) Deposit(amount Bitcoin){
	//fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}