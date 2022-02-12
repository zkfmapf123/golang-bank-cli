package main

import (
	"fmt"
	"golang-cli/src/cmd"
	"golang-cli/src/utils"
	"os"
)

var menu int

func main() {

	var loop = 0
	var isExit = false
	for loop < 1 {
		fmt.Println(utils.MENU_START)
		fmt.Print(utils.MENU_LIST)
		fmt.Scanf("%d", &menu)

		switch menu {
		case 1:
			cmd.Register()
		case 2:
			cmd.LookFor()
		case 3:
			cmd.AddMoney()
		case 4:
			cmd.MinMoney()
		case 5:
			cmd.PrintAll()
		default:
			fmt.Printf("%d is not valid", menu)
			isExit = !isExit
		}

		if isExit {
			os.Exit(0)
		}
		cmd.Clear()
		continue
	}
}
