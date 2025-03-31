package main

import (
	"errors"
	"fmt"
)

type Ether int

type Wallet struct {
	balance Ether
}

// In Go, when you call a function or a method the arguments are copied!

// struct pointers are automatically dereferenced! (no need to get value at the address)

func (w *Wallet) Deposit(amount Ether) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Ether) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Ether {
	return w.balance
}

// implement the string method on Ether
func (e Ether) String() string {
	return fmt.Sprintf("%d ETH", e)
}
