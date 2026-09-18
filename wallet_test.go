package main

import (
	"errors"
	"testing"
)

func TestDeposit_Success(t *testing.T) {
	w := newWallet(100, ACTIVE)

	if err := w.Deposit(50); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Balance != 150 {
		t.Errorf("balance = %d, want 150", w.Balance)
	}
}

func TestDeposit_ZeroAmount(t *testing.T) {
	am := int64(100)
	w := newWallet(am, ACTIVE)
	err := w.Deposit(0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("err = %v, want ErrInvalidAmount", err)
	}
	if w.Balance != am {
		t.Errorf("balance = %d, want = %d", w.Balance, am)
	}
}

func TestDeposit_Frozen(t *testing.T) {
	w := newWallet(100, FROZEN)
	err := w.Deposit(100)
	if !errors.Is(err, ErrWalletFrozen) {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDeposit_NegativeAmount(t *testing.T) {
	am := int64(100)
	w := newWallet(am, ACTIVE)
	err := w.Deposit(-10)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("err = %v, want ErrInvalidAmount", err)
	}
	if w.Balance != am {
		t.Errorf("balance = %d, want = %d", w.Balance, am)
	}
}

func TestWithdraw_Success(t *testing.T) {
	am := int64(100)
	red := int64(40)
	w := newWallet(am, ACTIVE)
	if err := w.Withdraw(red); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Balance != am-red {
		t.Errorf("balance = %d, want = %d", w.Balance, am)
	}
}

func TestWithdraw_ExactBalance(t *testing.T) {
	am := int64(100)
	w := newWallet(am, ACTIVE)
	if err := w.Withdraw(am); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Balance != 0 {
		t.Errorf("balance = %d, want = %d", w.Balance, 0)
	}
}

func TestWithdraw_ZeroAmount(t *testing.T) {
	am := int64(100)
	w := newWallet(am, ACTIVE)
	err := w.Withdraw(0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("err = %v, want ErrInvalidAmount", err)
	}
	if w.Balance != am {
		t.Errorf("balance = %d, want = %d", w.Balance, am)
	}
}

func TestWithdraw_Frozen(t *testing.T) {
	w := newWallet(100, FROZEN)
	err := w.Withdraw(10)
	if !errors.Is(err, ErrWalletFrozen) {
		t.Errorf("err = %d, ErrWalletFrozen", err)
	}
}

func TestWithdraw_InsufficientBalance(t *testing.T) {
	am := int64(100)
	req := int64(150)
	w := newWallet(am, ACTIVE)
	err := w.Withdraw(req)
	var e InsufficientBalanceError
	if !errors.As(err, &e) {
		t.Errorf("err = %v, want InsufficientBalanceError", err)
	}
	if e.Available != am {
		t.Errorf("balance = %d, want = %d", e.Available, am)
	}
	if e.Required != req {
		t.Errorf("required = %d, want = %d", e.Required, req)
	}
}

func TestFreeze_Success(t *testing.T) {
	w := newWallet(100, ACTIVE)
	if err := w.Freeze(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Status != FROZEN {
		t.Errorf("status = %v, want frozen", w.Status)
	}
}

func TestFreeze_Frozen(t *testing.T) {
	w := newWallet(100, FROZEN)
	err := w.Freeze()
	if !errors.Is(err, ErrWalletFrozen) {
		t.Fatalf("unexpected error: %v", err)
	}
}
