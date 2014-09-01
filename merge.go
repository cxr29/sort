package sort

import "sort"

func Merge(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	mergeSort(a, 0, n)
}

func LoopMerge(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	for i := 1; i < n; i <<= 1 {
		for j := 0; j < n-i; j += i << 1 {
			r := j + i<<1
			if r > n {
				r = n
			}
			if x := r - j; x == 2 && a.Less(j+1, j) {
				a.Swap(j, j+1)
			} else if x > 2 {
				merge(a, j, j+i, r)
			}
		}
	}
}

func mergeSort(a sort.Interface, l, r int) {
	n := r - l
	if n <= 2 {
		if n == 2 && a.Less(l+1, l) {
			a.Swap(l, l+1)
		}
		return
	}

	m := n>>1 + l

	mergeSort(a, l, m)
	mergeSort(a, m, r)

	merge(a, l, m, r)
}

func merge(a sort.Interface, l, m, r int) {
	// l: left
	// m: middle
	// n: next middle
	// r: right
	for n := m; l < m && n < r; l, m = l+(n-m), n {
		for l < n {
			if a.Less(n, l) {
				break
			} else {
				l++
			}
		}

		n++

		for n < r {
			if a.Less(n, l) {
				n++
			} else {
				break
			}
		}

		for x, y, z := l, m, n; ; {
			if n1, n2 := y-x, z-y; n1 > 0 && n2 > 0 {
				if n1 <= n2 {
					mergeSwap(a, x, y, n1)
					x += n1
					y += n1
				} else {
					mergeSwap(a, x+(n1-n2), y, n2)
					y -= n2
					z -= n2
				}
			} else {
				break
			}
		}

		//mergeReverse(a, l, m)
		//mergeReverse(a, m, n)
		//mergeReverse(a, l, n)
	}
}

func mergeSwap(a sort.Interface, i, j, n int) {
	for k := 0; k < n; k++ {
		a.Swap(i+k, j+k)
	}
}

func mergeReverse(a sort.Interface, l, r int) {
	for i, j := l, r-1; i < j; i, j = i+1, j-1 {
		a.Swap(i, j)
	}
}
