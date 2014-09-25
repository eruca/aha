package main

import (
	"fmt"
)

type QuickSort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data QuickSort) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data QuickSort, left, right int) {
	if left >= right {
		return
	}

	i := left
	j := right

	for i != j {
		for data.Less(left, j) && i < j {
			j--
		}

		for !data.Less(left, i) && i < j {
			i++
		}

		if i < j {
			data.Swap(i, j)
		}
	}

	data.Swap(left, i)

	quickSort(data, left, i-1)
	quickSort(data, i+1, right)
}

type Aint []int

func (a Aint) Len() int {
	return len(a)
}

func (a Aint) Less(i, j int) bool {
	if i == j {
		return false
	}
	return a[i] < a[j]
}

func (a Aint) Swap(i, j int) {
	if i == j {
		return
	}
	a[i], a[j] = a[j], a[i]
}

func main() {
	aint := []int{8, 23, 1, 4, 9, 32, 2, 72}

	Sort(Aint(aint))
	fmt.Println(aint)
}
