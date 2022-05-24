package main

import (
	"fmt"
	"strconv"
)

func setPalindrome(number int) bool {
	var strNum = strconv.Itoa(number)
	for i := 0; i < len(strNum); i++ {
		j := len(strNum)-1-i
		if strNum[i] != strNum[j] {
			return false
		}
	}

	return true
}

func main() {
	var num int

	fmt.Scan(&num)

	result := setPalindrome(num);

	if result {
		fmt.Println("this is palindrome")
	} else {
		fmt.Println("this is not palindrome")
	}
}