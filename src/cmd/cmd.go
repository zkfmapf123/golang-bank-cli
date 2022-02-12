package cmd

import (
	"bufio"
	"fmt"
	"golang-cli/src/banking"
	"golang-cli/src/utils"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Exit() {
	fmt.Println("다음으로 진행하려면 q를 입력하세요")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			return
		}
	}
}

func Register() {
	Clear()
	var name string
	var maxLimit int

	fmt.Printf(utils.REGISTER_NAME)
	fmt.Scanf("%s", &name)
	fmt.Printf(utils.REGISTER_MAX)
	fmt.Scanf("%d", &maxLimit)
	banking.Add(name, maxLimit)
}

func LookFor() {
	Clear()
	var name string
	fmt.Printf(utils.SEARCH_NAME)
	fmt.Scanf("%s", &name)
	ac, _, err := banking.Search(name)

	if err != nil {
		fmt.Println(err)
		Exit()
	}

	fmt.Println(ac)
	Exit()
}

func AddMoney() {
	Clear()
	var name string
	var addMoney int

	fmt.Printf(utils.DEPOSIT_NAME)
	fmt.Scanf("%s", &name)
	fmt.Printf(utils.DEPOSIT_BALANCE)
	fmt.Scanf("%d", &addMoney)

	err := banking.Deposit(name, addMoney)

	if err != nil {
		fmt.Println(err)
		Exit()
	}
}

func MinMoney() {
	Clear()
	var name string
	var minMoney int

	fmt.Printf(utils.WITHDRAW_NAME)
	fmt.Scanf("%s", &name)
	fmt.Printf(utils.WITHDRAW_BLANACE)
	fmt.Scanf("%d", &minMoney)

	err := banking.WithDraw(name, minMoney)

	if err != nil {
		fmt.Println(err)
		Exit()
	}

}

func PrintAll() {
	Clear()
	banking.PrintAll()
	Exit()
}
