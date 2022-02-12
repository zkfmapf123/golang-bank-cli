package banking

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	name     string
	balance  int
	maxLimit int
}

var accounts []BankAccount

func Add(name string, maxLimit int) {
	bc := BankAccount{name: name, balance: 0, maxLimit: maxLimit}
	accounts = append(accounts, bc)
	fmt.Println(accounts)
}

func Search(name string) (*BankAccount, int, error) {

	for index, value := range accounts {
		if value.name == name {
			return &value, index, nil
		}
	}

	return nil, 0, errors.New("not exits User : " + name)
}

// todo
func PrintAll() {
	for index, value := range accounts {
		fmt.Printf("%d : name : %s bal : %d, max : %d\n", index+1, value.name, value.balance, value.maxLimit)
	}
}

func Deposit(name string, deposit int) error {

	ac, idx, err := Search(name)

	if err != nil {
		return err
	}

	if ac.balance+deposit > ac.maxLimit {
		return errors.New("limit excess : fix to max")
	}

	accounts[idx].balance += deposit
	return nil
}

func WithDraw(name string, withdraw int) error {

	ac, idx, err := Search(name)

	if err != nil {
		return nil
	}

	if ac.balance-withdraw < 0 {
		return errors.New("balance < withdraw")
	}

	accounts[idx].balance -= withdraw
	return nil
}
