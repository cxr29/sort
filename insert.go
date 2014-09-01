package sort

import "sort"

func Insert(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		x := i
		for j := x - 1; j >= 0; j-- {
			if a.Less(x, j) {
				a.Swap(x, j)
				x = j
			}
		}
	}
}
