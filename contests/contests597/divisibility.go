package main

import (
	"fmt"
)

func main() {
	var k int64
	var a int64
	var b int64

	fmt.Scanf("%d", &k)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	//fmt.Println(k, a, b)
	var firstDiv int64
	if a == 0 || a % k == 0 {
		firstDiv = a
	} else {
		if a < 0 {
			firstDiv = a - a % k
		} else {
			firstDiv = a + (k - a % k)
		}
	}

	var lastDiv int64
	if b == 0 || b % k == 0 {
		lastDiv = b
	} else {
		if b < 0 {
			lastDiv = b - (k + b % k)
		} else {
			lastDiv = b - b % k
		}
	}

	//fmt.Println("firstDiv", firstDiv)
	//fmt.Println("lastDiv", lastDiv)
	countDivs := (lastDiv - firstDiv) / k + 1
	fmt.Println(countDivs)
}
