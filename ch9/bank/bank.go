package bank

import "sync"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

var (
	mu 		sync.RWMutex
	balance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

func DepositMutex(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func BalanceMutex() int {
	mu.RLock()
	defer mu.RUnlock()
	b := balance
	return b
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }