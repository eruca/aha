package main

import (
	"fmt"
)

var (
	n    int
	a    []int
	book []int
)

func dfs(step int) {
	if step > n {
		for i := 1; i <= n; i++ {
			fmt.Print(a[i])
		}
		fmt.Println("")
		return
	}
	for i := 1; i <= n; i++ {
		if book[i] == 0 {
			a[step] = i
			book[i] = 1

			dfs(step + 1)
			book[i] = 0
		}
	}
}

func main() {
	fmt.Println("请输入n的值")
	fmt.Scanf("%d", &n)
	fmt.Println("n:", n)
	a = make([]int, n+1)
	book = make([]int, n+1)
	fmt.Println(a)

	dfs(1)
}
