package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var k int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)

	studentsRating := make(map[int]int)
	for i := 1; i <= n; i++ {
		var rating int
		fmt.Scanf("%d", &rating)
		studentsRating[rating] = i
	}

	distinctRating := len(studentsRating)
	//fmt.Println(studentsRating)
	//fmt.Println(distinctRating)

	if distinctRating < k {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")

		diverseTeam := make([]int, k)
		diverseTeamIdx := 0
		for _, i := range studentsRating {
			diverseTeam[diverseTeamIdx] = i
			diverseTeamIdx++
		}

		sort.Ints(diverseTeam)
		for _, i := range diverseTeam {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
