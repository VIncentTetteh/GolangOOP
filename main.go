package main

import (
	"OOP/models"
	"fmt"
)

func main() {
	newUser := &models.User{"Vincent", "Accra", 43}
	wf := models.CreateWellsForgo()
	bitcoin := models.NewBitCoin(newUser)

	wf.Deposit(300)
	if err := wf.Redraw(30); err != nil {
		panic(err)
	}

	userInfo := bitcoin.Deposit(300)
	if berr := bitcoin.Redraw(250); berr != nil {
		panic(berr)
	}

	fmt.Printf("WF Current balance is %v", wf.GetBalance())
	fmt.Printf("BT Current balance is %v", bitcoin.GetBalance())
	fmt.Printf("BT user info is %v", userInfo)

}
