# Money transfer system

A command-line program in Go. Users keep a balance, and a payment system moves money from one user to another. Three workers process the queue at the same time, and each balance update is protected by a mutex.

## How to run

[Go](https://go.dev/dl/) 1.26.4 or later is required, as declared in `go.mod`. The workers are started with `wg.Add(1)` and `go ps.Worker(ch, &wg)`.

From the project directory:

```bash
go run main.go
```

To run the program with the race detector:

```bash
go run -race ./...
```

The program registers three users, queues three transfers, and lets three workers process them. Lines that start with `Transaction completed` can appear in any order. The balances for this data are:

- User1 — 2200
- User2 — 1300
- User3 — 2500

## What v4 adds

Builds on v3. This version makes balance changes safe when several workers touch the same user.

- `User` has a `sync.Mutex`.
- `Deposit` and `Withdraw` lock that mutex for the whole balance update, then unlock it.
