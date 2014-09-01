package sort

import (
	"math/rand"
	"sort"
)

var Rand *rand.Rand

func Quick(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	quick(a, 0, n)
}

func quick(a sort.Interface, l, r int) {
	n := r - l
	if n <= 2 {
		if n == 2 && a.Less(l+1, l) {
			a.Swap(l, l+1)
		}
		return
	}

	if Rand != nil {
		n = Rand.Int()%n + l
		a.Swap(l, n)
	}

	n = l
	for i, j := l+1, r-1; i <= j; i, j = i+1, j-1 {
		for i <= j {
			if a.Less(j, n) {
				a.Swap(n, j)
				n = j
				break
			} else {
				j--
			}
		}

		for i <= j {
			if a.Less(n, i) {
				a.Swap(i, n)
				n = i
				break
			} else {
				i++
			}
		}
	}

	quick(a, l, n)
	quick(a, n+1, r)
}
