# Q1

- Variant A returns ErrWalletFrozen
- Variant B returns ErrInvalidAmount

Variant A is better, but it pretty much depends if we want to do some extra expensive work (like send query to the DB) or not. State check should preceed input checks and the state check invariant can make operation illegal regardless of input. In this situation, user can interpet that if he fixes the amount, the program will accept their input (which is misleading), so ErrWalletFrozen is more informative.

# Q2

Required should be 150.

Messages:
1. Reguired=150: "Insufficient balance: required 150, available 100."
2. Required=50: "Insufficient balance: required 50 more, available 100."

Required=150 is better for UX because it matches the number they entered. There's no extra math or mental mapping needed - the user can immidiately connect the error to their own input.

# Q3

The main problem is that fmt.Errorf breaks errors.Is for API / handler layers for status codes / retries / etc. and domain errors are meaningful and complete. Usecase is orchestrator and should not be responsible for error transformation (for that it's better to create additional errors with more context if needed).