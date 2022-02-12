package main

import (
	"fmt"
	"golang-cli/src/utils"
	"os"
	"os/exec"
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
			fmt.Print(utils.REGISTER)
		case 2:
			fmt.Print(utils.SEARCH)
		case 3:
			fmt.Print(utils.DEPOSIT)
		case 4:
			fmt.Print(utils.WITHDRAW)
		case 5:
			fmt.Println(utils.DELETE)
		default:
			fmt.Printf("%d is not valid", menu)
			isExit = !isExit
		}

		if isExit {
			os.Exit(0)
		}

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		continue
	}

}
