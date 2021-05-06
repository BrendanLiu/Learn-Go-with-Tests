package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	got := wallet.Balance()
	want := BitCoin(20)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
