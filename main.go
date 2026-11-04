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

func printBalance(user *User) {
	fmt.Printf("%s has balance %f\n", user.ID, user.Balance)
}

func main() {
	user1 := &User{ID: "User1", Name: "Peter", Balance: 1000}
	user2 := &User{ID: "User2", Name: "Sarah", Balance: 2000}
	user3 := &User{ID: "User3", Name: "Michael", Balance: 3000}

	user1.Deposit(230.26)
	printBalance(user1)
	if err := user1.Withdraw(100); err != nil {
		fmt.Println(err)
	}
	printBalance(user1)

	user2.Deposit(300)
	printBalance(user2)
	user2.Deposit(102)
	printBalance(user2)
	if err := user2.Withdraw(50.5); err != nil {
		fmt.Println(err)
	}
	printBalance(user2)

	user3.Deposit(25)
	printBalance(user3)
	if err := user3.Withdraw(1100); err != nil {
		fmt.Println(err)
	}
	printBalance(user3)
	if err := user3.Withdraw(300.55); err != nil {
		fmt.Println(err)
	}
	printBalance(user3)
}
