package main

import (
	"errors"
	"fmt"
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
		return errors.New("insufficient funds")
	}

	u.Balance -= amount
	return nil
}

func main() {
	user1 := &User{"User1", "Peter", 1000}
	user2 := &User{"User2", "Sarah", 2000}
	user3 := &User{"User3", "Michael", 3000}

	user1.Deposit(230.26)
	user1.Withdraw(100)

	user2.Deposit(300)
	user2.Deposit(102)
	user2.Withdraw(50.5)

	user3.Deposit(25)
	user3.Withdraw(1100)
	user3.Withdraw(300.55)

	for _, u := range [3]*User{user1, user2, user3} {
		fmt.Printf("%s balance: %f\n", u.Name, u.Balance)
	}
}
