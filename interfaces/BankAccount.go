package interfaces

type BankAccount interface {
	Deposit(amount int)
	Redraw(amount int) error
	GetBalance() int
}
