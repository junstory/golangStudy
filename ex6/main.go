package main

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance int
}

type Exceptions struct {
	Msg string
}

func Deposit(account *BankAccount, amount int) {
	defer func() {
		fmt.Println(recover())
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Deposit Complete : ", account.Balance)
		}
	}()
	if amount <= 0 {
		panic("Can't deposit non-positive amount")
	}
	account.Balance += amount
}

func Withdraw(account *BankAccount, amount int) {
	defer func() {
		fmt.Println(recover())
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Withdraw Complete : ", account.Balance)
		}
	}()
	if amount > account.Balance {
		panic("Can't withdraw more than balance")
	}
	account.Balance -= amount
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != nil {
			fmt.Println("0으로 나눌 수 없습니다.")
		}
	}()
	user := BankAccount{"Tom", 100}
	Deposit(&user, -100)
	Withdraw(&user, 10000)
	test := 0
	fmt.Println(1 / test)
}
