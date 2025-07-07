package main

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("cannot Withdraw due to insufficient funds")

type Stringer interface {
	String() string
}
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		fmt.Println("I passed")
		// return errors.New("cannot Withdraw due to insufficient funds")
		return ErrorInsufficientFunds // Making sure test doesn't fail due to incorrect working
	}
	w.balance -= amount
	return nil
}

//Export -> Uppercase (Functions includes testfiles and everywhere)
//UnExported ->LowerCase (Functions exlusive to  testfiles)
