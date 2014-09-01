package sort

import (
	"math"
	"sort"
)

func Shell(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	for x := n / 2; x >= 1; x /= 2 {
		for i := x; i < n; i++ {
			for j := i - x; j >= 0; j -= x {
				if a.Less(j+x, j) {
					a.Swap(j+x, j)
				}
			}
		}
	}
}

func HibbardShell(a sort.Interface) {
	n := a.Len()
	if n < 2 {
		return
	}

	// 2^i - 1
	var x int
	for i := 1; ; i++ {
		if j := 1<<uint(i) - 1; j >= n || j <= x { // overflow
			break
		} else {
			x = j
		}
	}

	for ; x >= 1; x = (x+1)>>1 - 1 {
		for i := x; i < n; i++ {
			for j := i - x; j >= 0; j -= x {
				if a.Less(j+x, j) {
					a.Swap(j+x, j)
				}
			}
		}
	}
}

func KnuthShell(a sort.Interface) {
	incShell(a, knuthIncSeq)
}

func SedgewickShell(a sort.Interface) {
	incShell(a, sedgewickIncSeq)
}

func FibonacciShell(a sort.Interface) {
	incShell(a, fibonacciIncSeq)
}

func incShell(a sort.Interface, b []int) {
	n := a.Len()
	if n < 2 {
		return
	}

	for _, i := range b {
		if i >= n {
			continue
		}

		for j := i; j < n; j++ {
			for k := j - i; k >= 0; k -= i {
				if a.Less(k+i, k) {
					a.Swap(k+i, k)
				}
			}
		}
	}
}

var (
	knuthIncSeq     = knuthInc()
	sedgewickIncSeq = sedgewickInc()
	fibonacciIncSeq = fibonacciInc()
)

func knuthInc() []int {
	a := []int{1}
	for i := 2; ; i++ {
		j := int((math.Pow(3, float64(i)) - 1) / 2)

		if j <= a[len(a)-1] { // overflow
			break
		}

		a = append(a, j)
	}
	shellReverse(a)
	return a
}

func sedgewickInc() []int {
	a := []int{1}
	for i := 1; ; i++ {
		var j int

		if k := float64(i / 2); i%2 == 0 {
			j = int(9*(math.Pow(4, k)-math.Pow(2, k)) + 1)
		} else {
			j = int(math.Pow(2, k+2)*(math.Pow(2, k+2)-3) + 1)
		}

		if j <= a[len(a)-1] { // overflow
			break
		}

		a = append(a, j)
	}
	shellReverse(a)
	return a
}

func fibonacciInc() []int {
	x := 4 / (math.Sqrt(5) - 1)
	a := []int{0, 1}
	for i := 2; ; i++ {
		j := i - 2
		a = append(a, a[i-1]+a[j])

		a[j] = int(math.Pow(float64(a[i]), x))

		if j > 0 && a[j] < a[j-1] { // overflow
			a = a[:j]
			break
		}
	}
	shellReverse(a)
	return a
}

func shellReverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
