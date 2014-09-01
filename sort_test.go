package sort

import (
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
)

type SortFunc func(sort.Interface)

func (f SortFunc) Name() string {
	s := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	if i := strings.LastIndex(s, "."); i != -1 {
		return s[i+1:]
	}
	return s
}

func TestSort(t *testing.T) {
	testSort(t, Bubble)
	testSort(t, Select)
	testSort(t, Insert)
	testSort(t, Shell)
	testSort(t, HibbardShell)
	testSort(t, KnuthShell)
	testSort(t, SedgewickShell)
	testSort(t, FibonacciShell)
	testSort(t, Merge)
	testSort(t, Quick)
	testSort(t, Heap)
}

func testSort(t *testing.T, f SortFunc) {
	a := Ints{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	f(a)
	if !sort.IsSorted(a) {
		t.Fatalf("%s: Ints not sorted: %v\n", f.Name(), a)
	}

	b := Float64s{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
	f(b)
	if !sort.IsSorted(b) {
		t.Fatalf("%s: Float64s not sorted: %v\n", f.Name(), b)
	}

	c := Strings{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
	f(c)
	if !sort.IsSorted(c) {
		t.Fatalf("%s: Strings not sorted: %v\n", f.Name(), c)
	}
}

func BenchmarkFibonacciShellInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		FibonacciShell(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("FibonacciShell: Ints not sorted\n")
		}
	}
}

func BenchmarkFibonacciShellInt64K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		FibonacciShell(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("FibonacciShell: Ints not sorted\n")
		}
	}
}

func BenchmarkMergeInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		Merge(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("Merge: Ints not sorted\n")
		}
	}
}

func BenchmarkMergeInt64K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		Merge(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("Merge: Ints not sorted\n")
		}
	}
}

func BenchmarkMergeLoopInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		MergeLoop(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("MergeLoop: Ints not sorted\n")
		}
	}
}

func BenchmarkMergeLoopInt64K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		MergeLoop(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("MergeLoop: Ints not sorted\n")
		}
	}
}

func BenchmarkQuickInt1K(b *testing.B) {
	b.StopTimer()
	Rand = nil
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		Quick(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("Quick: Ints not sorted\n")
		}
	}
}

func BenchmarkQuickInt64K(b *testing.B) {
	b.StopTimer()
	Rand = nil
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		Quick(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("Quick: Ints not sorted\n")
		}
	}
}

func BenchmarkRandQuickInt1K(b *testing.B) {
	b.StopTimer()
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		Quick(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("RandQuick: Ints not sorted\n")
		}
	}
}

func BenchmarkRandQuickInt64K(b *testing.B) {
	b.StopTimer()
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		Quick(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("RandQuick: Ints not sorted\n")
		}
	}
}

func BenchmarkHeapInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		Heap(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("Heap: Ints not sorted\n")
		}
	}
}

func BenchmarkHeapInt64K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make(Ints, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		Heap(data)
		b.StopTimer()
		if i == 0 && !sort.IsSorted(data) {
			b.Fatalf("Heap: Ints not sorted\n")
		}
	}
}
