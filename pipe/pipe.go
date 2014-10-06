package main

import (
	"fmt"
)

var (
	a          [51][51]int
	book       [51][51]int
	n, m, flag int = 5, 4, 0
)

func dfs(x, y, front int) {
	if x == n && y == m+1 {
		flag = 1
		return
	}

	if x < 1 || x > n || y < 1 || y > m {
		return
	}
	if book[x][y] == 1 {
		return
	}

	book[x][y] = 1

	if a[x][y] == 5 || a[x][y] == 6 {
		switch front {
		case 1:
			dfs(x, y+1, 1)
		case 2:
			dfs(x+1, y, 2)
		case 3:
			dfs(x, y-1, 3)
		case 4:
			dfs(x-1, y, 4)
		}
	}
	if a[x][y] >= 1 && a[x][y] <= 4 {
		switch front {
		case 1:
			dfs(x+1, y, 2)
			dfs(x-1, y, 4)
		case 2:
			dfs(x, y+1, 1)
			dfs(x, y-1, 3)
		case 3:
			dfs(x-1, y, 4)
			dfs(x+1, y, 2)
		case 4:
			dfs(x, y+1, 1)
			dfs(x, y-1, 3)
		}
	}
	book[x][y] = 0
}

func main() {
	var res = []int{5, 3, 5, 3, 1, 5, 3, 0, 2, 3, 5, 1, 6, 1, 1, 5, 1, 5, 5, 4}
	var index int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] = res[index]
			index++
		}
	}
	dfs(1, 1, 1)
	if flag == 0 {
		fmt.Println("impossible")
	} else {
		fmt.Println("找到铺设方法")
	}
}
