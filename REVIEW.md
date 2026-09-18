# Buggy Code - Fix

The problem is that errors.Is diretly compares 2 errors. In out situation it will compare Code and Message fields. If they're matching (both fields), it will return true, otherwise false.

To fix a test we need to implement Is method for WalletError.

```go
func(e WalletError) Is(target error) bool {
    t, ok := target.(WalletError)
    if !ok {
        return false
    }
    return e.Code == t.Code
}

func TestWithdraw_InsufficientBalance(t *testing.T) {
    w := NewWallet("w1", "owner1", 100)
    err := w.Withdraw(500)
    if !errors.Is(err, ErrInsufficientBalance) {
        t.Fatal("expected ErrInsufficientBalance")  
    }
}
```
