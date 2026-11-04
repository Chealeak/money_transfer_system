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

## Versions

Each branch adds the next part of the program. This branch is v4.

### v1

Users and balances.

- `User` has an ID, a name, and a balance.
- `Deposit` adds to the balance.
- `Withdraw` subtracts from the balance and returns `ErrInsufficientFunds` when the balance is too low.
- `main` runs a few deposits and withdrawals and prints the updated balance after each one.

### v2

Transfers through a queue, processed one after another.

- `Transaction` holds the sender ID, the receiver ID, and the amount.
- `PaymentSystem` stores users in a map and keeps a `TransactionQueue`.
- `AddUser` registers a user. `AddTransaction` appends a transfer to the queue.
- `ProcessingTransactions` withdraws the amount from the sender and deposits it to the receiver. It returns `ErrSenderNotFound` or `ErrReceiverNotFound` when a participant is missing, and `ErrInsufficientFunds` when the sender's balance is too low.
- `main` fills the queue and processes it in order. `Transaction completed` is printed only when `ProcessingTransactions` returns no error.

### v3

The same queue, processed by three workers at the same time.

- `Worker` reads transactions from a channel, runs `ProcessingTransactions` for each one, and calls `defer wg.Done()` when it returns.
- A failed transfer is printed and does not stop the other transfers.
- `main` calls `wg.Add(1)` and then `go ps.Worker(ch, &wg)` for each worker, sends the queue into a buffered channel, closes the channel, and waits with `wg.Wait()`.

### v4

Safe balance updates while those workers run, plus a Go module.

- `User` has a `sync.Mutex`.
- `Deposit` and `Withdraw` lock that mutex for the whole balance update, then unlock it.
