package main

import (
	"fmt"
	"os"
)

var GREETING = "Back Account CLI"
var MENU = "1. 은행계좌 등록\n2. 은행계좌 검색\n3. 입금\n4. 출금\n5. 은행계좌 삭제\n"
var SELECT_MENU = "select menu : "

var menu int

func main() {

	// 해당 문자열 입력
	fmt.Println(GREETING)
	fmt.Print(MENU)
	fmt.Print(SELECT_MENU)
	fmt.Scanf("%d", &menu)

	switch menu {
	case 1:
		fmt.Print("은행계좌 등록")
	case 2:
		fmt.Print("은행계좌 검색")
	case 3:
		fmt.Print("은행계좌 입금")
	case 4:
		fmt.Print("은행계좌 출금")
	default:
		fmt.Printf("%d is not valid", menu)
		os.Exit(0)
	}

}
