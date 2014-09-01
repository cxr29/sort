package sort

type Ints []int

func (a Ints) Len() int {
	return len(a)
}

func (a Ints) Less(i, j int) bool {
	if a[i] > a[j] {
		return false
	} else if a[i] < a[j] {
		return true
	}
	return i < j
}

func (a Ints) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Float64s []float64

func (a Float64s) Len() int {
	return len(a)
}

func (a Float64s) Less(i, j int) bool {
	if a[i] > a[j] {
		return false
	} else if a[i] < a[j] {
		return true
	}
	return i < j
}

func (a Float64s) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Strings []string

func (a Strings) Len() int {
	return len(a)
}

func (a Strings) Less(i, j int) bool {
	if a[i] > a[j] {
		return false
	} else if a[i] < a[j] {
		return true
	}
	return i < j
}

func (a Strings) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
