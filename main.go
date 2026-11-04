package main

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrSenderNotFound    = errors.New("sender not found")
	ErrReceiverNotFound  = errors.New("receiver not found")
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) error {
	if u.Balance < amount {
		return ErrInsufficientFunds
	}

	u.Balance -= amount
	return nil
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users            map[string]*User
	TransactionQueue []Transaction
}

func (ps *PaymentSystem) AddUser(user *User) {
	ps.Users[user.ID] = user
}

func (ps *PaymentSystem) AddTransaction(transaction Transaction) {
	ps.TransactionQueue = append(ps.TransactionQueue, transaction)
}

func (ps *PaymentSystem) ProcessingTransactions(transaction Transaction) error {
	sender, ok := ps.Users[transaction.FromID]
	if !ok {
		return ErrSenderNotFound
	}

	receiver, ok := ps.Users[transaction.ToID]
	if !ok {
		return ErrReceiverNotFound
	}

	err := sender.Withdraw(transaction.Amount)
	if err != nil {
		return err
	}

	receiver.Deposit(transaction.Amount)
	return nil
}

func main() {
	ps := &PaymentSystem{Users: make(map[string]*User)}

	user1 := &User{ID: "User1", Name: "Peter", Balance: 1000}
	user2 := &User{ID: "User2", Name: "Sarah", Balance: 2000}
	user3 := &User{ID: "User3", Name: "Michael", Balance: 3000}

	users := []*User{user1, user2, user3}
	for i := range users {
		ps.AddUser(users[i])
		fmt.Printf("Added user %s with balance %f\n", users[i].ID, users[i].Balance)
	}

	transactions := []Transaction{
		{FromID: user1.ID, ToID: user2.ID, Amount: 300},
		{FromID: user2.ID, ToID: user3.ID, Amount: 1000},
		{FromID: user3.ID, ToID: user1.ID, Amount: 1500},
	}
	for i := range transactions {
		ps.AddTransaction(transactions[i])
		fmt.Printf("Transaction added to queue: %s sent %f to %s\n", transactions[i].FromID, transactions[i].Amount, transactions[i].ToID)
	}

	for _, t := range ps.TransactionQueue {
		err := ps.ProcessingTransactions(t)
		if err != nil {
			fmt.Println("An error occurred during the transaction:", err)
		} else {
			fmt.Printf("Transaction completed: %s sent %f to %s\n", t.FromID, t.Amount, t.ToID)
		}
	}

	fmt.Println("Result:")
	for _, user := range users {
		fmt.Printf("%s has balance %f\n", user.ID, user.Balance)
	}
}
