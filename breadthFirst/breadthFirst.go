package main

import (
	"fmt"
)

type node struct {
	x int
	y int
	f int
	s int
}

func main() {
	var que [2501]node
	var a, book [51][51]int
	var next = [4][2]int{{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var head, tail int

	var n, m int
	fmt.Println("input the n×m")
	fmt.Scanf("%d %d", &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	var startX, startY, p, q int
	fmt.Scanf("%d %d %d %d", &startX, &startY, &p, &q)

	head = 1
	tail = 1

	que[tail].x = startX
	que[tail].y = startY
	que[tail].f = 0
	que[tail].s = 0
	tail++
	book[startX][startY] = 1

	var flag = 0
	var tx, ty int
	for head < tail {
		for i := 0; i <= 3; i++ {
			tx = que[head].x + next[i][0]
			ty = que[head].y + next[i][1]

			if tx < 1 || tx > n || ty < 1 || ty > m {
				continue
			}

			if a[tx][ty] == 0 && book[tx][ty] == 0 {
				book[tx][ty] = 1
				que[tail].x = tx
				que[tail].y = ty
				que[tail].f = head
				que[tail].s = que[head].s + 1
				tail++
			}

			if tx == p && ty == q {
				flag = 1
				break
			}
		}

		if flag == 1 {
			break
		}
		head++
	}

	fmt.Print(que[tail-1].s)
}
