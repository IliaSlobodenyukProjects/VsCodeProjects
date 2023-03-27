// горутины, одновременно работающие с общими данными, сами собой не могут синхронизироваться
package main

import (
	"fmt"
)

type Account struct {
	balance float64
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	fmt.Println("depositing: ", amount, '\n')
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
		fmt.Println("here")
		return
	}
	fmt.Println("withdrawing: ", amount, '\n')
	a.balance -= amount
}

func main() {
	acc := Account{}

	// Старт 10 горутин
	for i := 0; i < 10; i++ {
		go func() {
			// каждая из которых производит операции с аккаунтом
			for j := 0; j < 10; j++ {
				// Иногда снимает деньги
				if j%2 == 1 {
					acc.Withdraw(50)
					continue
				}
				// Иногда кладет
				acc.Deposit(50)
			}
		}()
	}
	fmt.Scanln()
	fmt.Println(acc.Balance())
}

func closure() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("GOT", i)
		}(i)
	}
}
