# Money transfer system

A command-line program in Go. This version models users and their balances: money can be deposited and withdrawn, and a withdrawal fails when the balance is too low.

## How to run

[Go](https://go.dev/dl/) is required.

From the project directory:

```bash
go run main.go
```

The program creates three users, applies a few deposits and withdrawals, and prints the updated balance after each one:

```text
User1 has balance 1230.260000
User1 has balance 1130.260000
User2 has balance 2300.000000
User2 has balance 2402.000000
User2 has balance 2351.500000
User3 has balance 3025.000000
User3 has balance 1925.000000
User3 has balance 1624.450000
```

## What v1 adds

This is the first version.

- `User` with an ID, a name, and a balance.
- `Deposit` adds to the balance.
- `Withdraw` subtracts from the balance and returns `ErrInsufficientFunds` when the balance is too low.
- `main` runs a few deposits and withdrawals and prints the balance after each one.
