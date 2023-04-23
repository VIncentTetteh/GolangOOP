package models

import "errors"

type WellsForgo struct {
	balance int
}

func CreateWellsForgo() *WellsForgo {
	return &WellsForgo{balance: 0}
}

func (w *WellsForgo) Deposit(amount int) {
	//TODO implement me
	w.balance += amount
}

func (w *WellsForgo) Redraw(amount int) error {
	//TODO implement me
	newBalance := w.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient fund")
	}
	w.balance = newBalance
	return nil
}

func (w *WellsForgo) GetBalance() int {
	//TODO implement me
	return w.balance
}
