package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, Bitcoin(10), wallet)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, Bitcoin(10), wallet)
	})
	t.Run("Withdraw when insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(200))
		assertBalance(t, startingBalance, wallet)
		assertError(t, err, ErrorInsufficientFunds.Error())
	})
}
func assertBalance(t testing.TB, want Bitcoin, wallet Wallet) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Got %s and want %s", got, want)
	}

}

// You can make want as error and you won't need to do ErrorInsufficientFunds.Error()
func assertError(t testing.TB, got error, want string) {
	t.Helper()
	// if err == nil {
	// 	t.Fatal("Wanted an Error but didn't get one")
	// }
	if got == nil {
		t.Fatal("didnt' get an error but wanted one")
	}
	if got.Error() != want { //Converting error to string with .Error()
		t.Errorf("Got %q and want %q", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't wanted one")

	}

}

//NOTE : What is the diff b/w t testing.TB & *testing.T and *testing.B
//1) Tb -> is common interface , shared b/w T and B and is used in Error Handling,logging and control
//2) T -> Is used for Unit Testing
//3) B-> Is used for Benchmark (Iteraions b.N etc)
//4) t.helper() is used for quick & ez debugging
