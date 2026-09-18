# Backend Developer Assessment - Junior Level (Task 2)

## Instructions

1. Complete all tasks below
2. Push your solution to a **public GitHub repository**
3. Reply to this email with your repository URL
4. Deadline: 48 hours from receiving this email

---

## Task: Wallet Entity with Error Handling

Create a `Wallet` domain entity for a payment system with proper error handling.

### Requirements

**Fields:**
- `id` - WalletID
- `ownerID` - OwnerID
- `balance` - in cents (int64)
- `status` - ACTIVE, FROZEN

**Methods:**
- `Deposit(amount int64) error`
- `Withdraw(amount int64) error`
- `Freeze() error`

**Business Rules:**
- Amount must be positive (> 0) for both Deposit and Withdraw
- Cannot operate on FROZEN wallet
- Cannot withdraw more than available balance

**Error Types Required:**

```go
var (
    ErrInsufficientBalance = errors.New("insufficient balance")
    ErrInvalidAmount       = errors.New("invalid amount")
    ErrWalletFrozen        = errors.New("wallet is frozen")
)

// Structured error - must work with errors.Is()
type InsufficientBalanceError struct {
    Required  int64
    Available int64
}
```

---

## Buggy Code - Find the Problem

This test is failing. The code compiles but `errors.Is` never matches. Explain why in `REVIEW.md`:

```go
// errors.go
type WalletError struct {
    Code    string
    Message string
}

func (e WalletError) Error() string {
    return e.Message
}

var ErrInsufficientBalance = WalletError{Code: "E001", Message: "insufficient balance"}

// wallet.go
func (w *Wallet) Withdraw(amount int64) error {
    if w.balance < amount {
        return WalletError{
            Code:    "E001",
            Message: fmt.Sprintf("need %d, have %d", amount, w.balance),
        }
    }
    w.balance -= amount
    return nil
}

// wallet_test.go
func TestWithdraw_InsufficientBalance(t *testing.T) {
    w := NewWallet("w1", "owner1", 100)
    err := w.Withdraw(500)

    if !errors.Is(err, ErrInsufficientBalance) {
        t.Fatal("expected ErrInsufficientBalance")  // ALWAYS FAILS
    }
}
```

Your `REVIEW.md` must include:
1. Why `errors.Is` returns false
2. The exact fix (show corrected code)

---

## Questions - Answer in ANSWERS.md

**Q1:** Look at these two implementations. For a FROZEN wallet with balance=100, calling `Withdraw(-50)`:

```go
// Version A
func (w *Wallet) Withdraw(amount int64) error {
    if w.status == StatusFrozen { return ErrWalletFrozen }
    if amount <= 0 { return ErrInvalidAmount }
    // ...
}

// Version B
func (w *Wallet) Withdraw(amount int64) error {
    if amount <= 0 { return ErrInvalidAmount }
    if w.status == StatusFrozen { return ErrWalletFrozen }
    // ...
}
```

- What does Version A return?
- What does Version B return?
- Which is correct? Explain your reasoning.

**Q2:** Your `InsufficientBalanceError` has `Required` and `Available` fields. If balance=100 and someone tries to withdraw 150:

- Should `Required` be `150` (the requested amount) or `50` (the deficit)?
- Write the user-facing error message for each choice
- Which is better UX?

**Q3:** Why should domain errors NOT be wrapped with `fmt.Errorf("failed: %w", err)` in the usecase layer?

---

## Repository Structure

```
your-repo/
├── wallet.go          # Your implementation
├── errors.go          # Error types
├── wallet_test.go     # Unit tests
├── REVIEW.md          # Bug analysis
└── ANSWERS.md         # Question answers
```

---

## Evaluation

Your submission will be evaluated against our engineering standards document. Key areas:
- Proper error types with Is() method implementation
- Validation order (inputs before state checks)
- Domain errors returned as-is, not wrapped
- Test coverage including errors.Is and errors.As usage
- Structured errors with context where needed
