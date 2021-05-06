package main

import "fmt"

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

// String() 方法让现有BitCoin类型添加一些领域内特定的功能
func (b BitCoin) String() string {
	return fmt.Sprintf("%d Btc", b)
}
