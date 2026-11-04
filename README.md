# Money transfer system

A command-line program in Go. Users keep a balance, and a payment system moves money from one user to another. Three workers process the queue at the same time, and each balance update is protected by a mutex.

## How to run

[Go](https://go.dev/dl/) 1.25 or later is required (`sync.WaitGroup.Go`).

From the project directory:

```bash
go run main.go
```

The program registers three users, queues three transfers, and lets three workers process them. Lines that start with `Transaction completed` can appear in any order. The balances for this data are:

- User1 — 2200
- User2 — 1300
- User3 — 2500

## What v4 adds

Builds on v3. This version makes balance changes safe when several workers touch the same user.

- `User` has a `sync.Mutex`.
- `Deposit` and `Withdraw` lock that mutex for the whole balance update, then unlock it.
