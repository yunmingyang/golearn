package main

import (
	"sync"
)



var (
	mu sync.Mutex
	balance int
)

func deposit(amount int) {
	balance += amount
}

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}
// func Deposit(amount int) {
// 	mu.Lock()
// 	balance = balance + amount
// 	mu.Unlock()
// }

func Balance() int {
	// mu.Lock()
	// b := balance
	// mu.Unlock()
	// return b
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// // Note: not atomic!
// func WithDraw(amount int) bool {
// 	// // Note: incorrect, as in Deposit, the lock will be repeated
// 	// mu.Lock()
// 	// defer mu.Unlock()
// 	Deposit(-amount)
// 	if balance < 0 {
// 		Deposit(amount)
// 		return false // insufficient funds
// 	}
// 	return true
// }

func WithDraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}

	return true
}