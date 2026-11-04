package main

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
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

func main() {
	user1 := &User{ID: "User1", Name: "Peter", Balance: 1000}
	user2 := &User{ID: "User2", Name: "Sarah", Balance: 2000}
	user3 := &User{ID: "User3", Name: "Michael", Balance: 3000}

	user1.Deposit(230.26)
	user1.Withdraw(100)

	user2.Deposit(300)
	user2.Deposit(102)
	user2.Withdraw(50.5)

	user3.Deposit(25)
	user3.Withdraw(1100)
	user3.Withdraw(300.55)

	users := []*User{user1, user2, user3}
	fmt.Println("Result:")
	for _, user := range users {
		fmt.Printf("%s has balance %f\n", user.ID, user.Balance)
	}
}
