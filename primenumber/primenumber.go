package main

import (
	"fmt"
)

func main() {
	primeNumber(255)
}

//verifyNumber for is verify number
func verifyNumber(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//primeNumber show number prime
func primeNumber(n int) {
	for i := 2; i <= n; i++ {
		if verifyNumber(i) {
			fmt.Println(i)
		}

	}
}
