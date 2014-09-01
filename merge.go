package sort

import "sort"

func Merge(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	merge(a, 0, n)
}

func merge(a sort.Interface, l, r int) {
	n := r - l
	if n <= 2 {
		if n == 2 && a.Less(l+1, l) {
			a.Swap(l, l+1)
		}
		return
	}

	m := n>>1 + l

	merge(a, l, m)
	merge(a, m, r)

	// l: left
	// m: middle
	// n: next middle
	// r: right
	for n = m; l < m && n < r; l, m = l+(n-m), n {
		for l < n {
			if a.Less(n, l) {
				break
			} else {
				l++
			}
		}

		for n < r {
			if a.Less(n, l) {
				n++
			} else {
				break
			}
		}

		mergeReverse(a, l, m)
		mergeReverse(a, m, n)
		mergeReverse(a, l, n)
	}
}

func mergeReverse(a sort.Interface, l, r int) {
	for i, j := l, r-1; i < j; i, j = i+1, j-1 {
		a.Swap(i, j)
	}
}
