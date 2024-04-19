package main

import (
	"fmt"
	"strings"
)

func main() {
	passwd := "MyRandomP2ss"
	if pwdChecker(passwd) {
		fmt.Println("Strong Password")
		return
	}
	fmt.Println("Weak password")

}

func pwdChecker(pass string) bool {
	if len(pass) < 9 {
		return false
	}
	ucCheck := strings.ContainsAny(pass, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lcCheck := strings.ContainsAny(pass, "abcdefghijklmnopqrstuvwxyz")
	numCheck := strings.ContainsAny(pass, "1234567890")

	return ucCheck && lcCheck && numCheck
}
