# Money transfer system

A command-line program in Go. Users keep a balance, and a payment system moves money from one user to another. This version processes the transaction queue with several workers at the same time.

## How to run

[Go](https://go.dev/dl/) 1.22 or later is required (`for range 3`).

From the project directory:

```bash
go run main.go
```

The program registers three users, queues three transfers, and lets three workers process them. Lines that start with `Transaction completed` can appear in any order. The balances for this data are:

- User1 — 2200
- User2 — 1300
- User3 — 2500

## What v3 adds

Builds on v2. Processing moves off the main goroutine.

- `Worker` reads transactions from a channel, runs `ProcessingTransactions` for each one, and calls `defer wg.Done()` when it returns.
- A failed transfer is printed and does not stop the other transfers.
- `main` calls `wg.Add(1)` and then `go ps.Worker(ch, &wg)` for each of the three workers, sends the queue into a buffered channel, closes the channel, and waits with `wg.Wait()`.
