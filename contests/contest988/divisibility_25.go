package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	processedString = make(map[string]bool)
	divisibleBy25   = make(map[string]int)
)

func main() {
	var inputInt int
	fmt.Scanf("%d", &inputInt)
	inputStr := strconv.Itoa(inputInt)
	inputSplit := strings.Split(inputStr, "")

	//fmt.Println(inputSplit)
	findDivisible(inputSplit, 0)
}

func findDivisible(s []string, swapCount int) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Printf("i: %d, len: %d\n", i, len(s)-1)
		toSwap := strings.Join(s[:], "")
		fmt.Printf("start swapping %s: %d to %d with count: %d\n", toSwap, i, i+1, swapCount)
		swapElements(swapCount, i, i+1, s)
		fmt.Printf("done swapping %s: %d to %d with count: %d\n", toSwap, i, i+1, swapCount)
	}
}

func swapElements(swapCount, from, to int, s []string) {
	//before := strings.Join(s[:], "")
	s[from], s[to] = s[to], s[from]
	after := strings.Join(s[:], "")

	//fmt.Printf("swapped indices %d and %d of %s, resulting %s\n", from, to, before, after)

	if processedString[after] {
		return
	}
	processedString[after] = true

	inProcessInt, _ := strconv.Atoi(after)
	if inProcessInt%25 == 0 {
		fmt.Printf("found one divisible by 25: %s, swaps: %d\n", after, swapCount)
		divisibleBy25[after] = swapCount
	}
	findDivisible(s, swapCount+1)
}
