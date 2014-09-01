package sort

import "sort"

func Heap(a sort.Interface) {
	heap(a, maxHeap)
}

// min, ..., max
func MaxHeap(a sort.Interface) {
	heap(a, maxHeap)
}

// max, ..., min
func MinHeap(a sort.Interface) {
	heap(a, minHeap)
}

func heap(a sort.Interface, f func(sort.Interface, int, int)) {
	n := a.Len()
	if n < 2 {
		return
	}

	for i := n/2 - 1; i >= 0; i-- {
		f(a, i, n)
	}

	for i := n - 1; i >= 1; i-- {
		a.Swap(0, i)
		f(a, 0, i)
	}
}

func maxHeap(a sort.Interface, i, n int) {
	for j := 2*i + 1; j < n; j = 2*i + 1 {
		if j+1 < n && a.Less(j, j+1) {
			j++
		}

		if a.Less(i, j) {
			a.Swap(i, j)
			i = j
		} else {
			break
		}
	}
}

func minHeap(a sort.Interface, i, n int) {
	for j := 2*i + 1; j < n; j = 2*i + 1 {
		if j+1 < n && a.Less(j+1, j) {
			j++
		}

		if a.Less(j, i) {
			a.Swap(j, i)
			i = j
		} else {
			break
		}
	}
}
