package main

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrWalletFrozen        = errors.New("wallet is frozen")
)

type InsufficientBalanceError struct {
	Required  int64
	Available int64
}

func (e InsufficientBalanceError) Error() string {
	return fmt.Sprintf("insufficient balance: have: %d, want: %d", e.Available, e.Required)
}

func (e InsufficientBalanceError) Is(target error) bool {
	t, ok := target.(InsufficientBalanceError)
	if !ok {
		return false
	}
	return e.Available == t.Available && e.Required == t.Required
}
