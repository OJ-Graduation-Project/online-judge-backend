/* Fenwick tree (https://en.wikipedia.org/wiki/Fenwick_tree) */
package fenwick

import "errors"

var IndexError = errors.New("Index is out of range")

// Basic structure
type Fenwick struct {
	data []int
}

// New creates a Fenwick tree of fixed size
func New(size uint) *Fenwick {
	var f Fenwick
	f.data = make([]int, size+1)
	return &f
}

func (f *Fenwick) Update(i, v int) error {
	if i < 0 || i >= len(f.data) {
		return IndexError
	}
	i++
	for ; i < len(f.data); i += (i & -i) {
		f.data[i] += v
	}
	return nil
}

// Return sum of value on segment [begin, end] (including both begin and end)
func (f *Fenwick) Sum(begin, end int) int {
	endSum := f.Query(end)
	beginSum := f.Query(begin - 1)
	return endSum - beginSum
}

func (f *Fenwick) Query(i int) int {
	i++
	if i >= len(f.data) {
		i = len(f.data) - 1
	}
	result := 0
	for ; i > 0; i -= (i & -i) {
		result += f.data[i]
	}
	return result
}

func (f *Fenwick) Suffix(i int) int {
	return f.Query(len(f.data)-2) - f.Query(i-1)
}

func (f *Fenwick) Find(k int) int {
	l := 0
	r := len(f.data) - 1

	for l < r {
		m := l + (r-l+1)/2
		v := f.Suffix(m)
		if v >= k {
			l = m
		} else {
			r = m-1
		}
	}
	return l
}

func(f *Fenwick) Before(i int) int {
	return f.Suffix(i+1)
}