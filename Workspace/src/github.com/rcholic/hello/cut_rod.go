package main

import (
	"fmt"
)

var r []int //slice

func max(a int, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

func memoized_cut_rod(price [8]int, n int) int {
	if n <= 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		r[i] = -10
	}

	return memoized_cut_rod_aux(price, n, r)
}

func memoized_cut_rod_aux(p [8]int, n int, rValue []int) int {

	if rValue[n] >= 0 {
		return rValue[n]
	}
	q := 0
	if n == 0 {
		q = 0
	} else {
		q = -10
		for i := 1; i < n; i++ {
			q = max(q, p[i]+memoized_cut_rod_aux(p, n-i, rValue))
		}
	}
	rValue[n] = q

	return q
}

func main() {

	price := [8]int{1, 5, 8, 9, 10, 17, 17, 20}
	size := len(price)
	fmt.Println(memoized_cut_rod(price, size))

}
