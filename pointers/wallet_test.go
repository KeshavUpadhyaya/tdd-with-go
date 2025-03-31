package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Ether) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error but didn't want one")
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		wallet.Deposit(Ether(10))

		assertBalance(t, wallet, Ether(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Ether(20)}

		err := wallet.Withdraw(Ether(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Ether(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Ether(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Ether(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}
