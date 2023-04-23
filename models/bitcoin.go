package models

import (
	"errors"
)

type Bitcoin struct {
	user    User
	balance int
	fee     int
}

func NewBitCoin(user *User) *Bitcoin {
	return &Bitcoin{
		balance: 0,
		fee:     300,
		user:    *user,
	}
}

func (b *Bitcoin) Deposit(amount int) *Bitcoin {
	//TODO implement me
	b.balance += amount
	return b
}

func (b *Bitcoin) Redraw(amount int) error {
	//TODO implement me
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient funds in BT account")
	}
	return nil
}

func (b *Bitcoin) GetBalance() int {
	//TODO implement me
	return b.balance
}
