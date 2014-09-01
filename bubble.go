package sort

import "sort"

func Bubble(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a.Less(j+1, j) {
				a.Swap(j+1, j)
			}
		}
	}
}
