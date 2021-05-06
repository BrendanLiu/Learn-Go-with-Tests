package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want BitCoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			// Fatal 如果被调用，将不会进行接下来的测试
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))

		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		err := wallet.Withdraw(BitCoin(10))

		assertBalance(t, wallet, BitCoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, BitCoin(20))
		assertError(t, err, InsufficientFundsError)
	})
}
