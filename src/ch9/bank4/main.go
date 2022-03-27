package main

import "sync"



var (
	mu sync.RWMutex
	balance int
)

func deposit(amount int) {
	balance += amount
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()

	return balance
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()

	deposit(amount)
}

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