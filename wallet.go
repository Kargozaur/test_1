package main

import "uuid"

type WalletStatus = string

const (
	ACTIVE WalletStatus = "active"
	FROZEN WalletStatus = "frozen"
)

type Wallet struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Balance int64
	Status  WalletStatus
}

func newWallet(balance int64, status WalletStatus) *Wallet {
	return &Wallet{
		ID:      uuid.NewV7(),
		OwnerID: uuid.NewV7(),
		Balance: balance,
		Status:  status,
	}
}
func (w *Wallet) Deposit(amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if w.isFrozen() {
		return ErrWalletFrozen
	}
	w.Balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if w.isFrozen() {
		return ErrWalletFrozen
	}
	if w.Balance < amount {
		return InsufficientBalanceError{Available: w.Balance, Required: amount}
	}
	w.Balance -= amount
	return nil
}

func (w *Wallet) Freeze() error {
	if w.isFrozen() {
		return ErrWalletFrozen
	}
	w.Status = FROZEN
	return nil
}

func (w *Wallet) isFrozen() bool {
	return w.Status == FROZEN
}
