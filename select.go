package sort

import "sort"

func Select(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	for i := 0; i < n; i++ {
		x := i
		for j := x + 1; j < n; j++ {
			if a.Less(j, x) {
				x = j
			}
		}
		if i != x {
			a.Swap(i, x)
		}
	}
}
