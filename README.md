# Money transfer system

A command-line program in Go. Users keep a balance, and a payment system moves money from one user to another through a queue of transactions.

## How to run

[Go](https://go.dev/dl/) is required.

From the project directory:

```bash
go run main.go
```

The program registers three users, queues three transfers, processes the queue in order, and prints the balances:

```text
Added user User1 with balance 1000.000000
Added user User2 with balance 2000.000000
Added user User3 with balance 3000.000000
Transaction added to queue: User1 sent 300.000000 to User2
Transaction added to queue: User2 sent 1000.000000 to User3
Transaction added to queue: User3 sent 1500.000000 to User1
Transaction completed: User1 sent 300.000000 to User2
Transaction completed: User2 sent 1000.000000 to User3
Transaction completed: User3 sent 1500.000000 to User1
Result:
User1 has balance 2200.000000
User2 has balance 1300.000000
User3 has balance 2500.000000
```

## What v2 adds

Builds on v1. This version adds transfers.

- `Transaction` holds the sender ID, the receiver ID, and the amount.
- `PaymentSystem` stores users in a map and keeps a `TransactionQueue`.
- `AddUser` registers a user. `AddTransaction` appends a transfer to the queue.
- `ProcessingTransactions` withdraws the amount from the sender and deposits it to the receiver.
- `ErrSenderNotFound` and `ErrReceiverNotFound` are returned when a participant is missing.
- `main` fills the queue and processes it one transaction after another.
