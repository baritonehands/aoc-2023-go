package utils

import (
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"iter"
	"slices"
	"strings"
)

type Pair struct {
	X, Y int
}

func (p Pair) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func Split2(s string) (string, string) {
	arr := strings.SplitN(s, " ", 2)
	return arr[0], arr[1]
}

func SeqSet[I iter.Seq[T], T comparable](iter I) map[T]bool {
	return it.Fold(iter, func(m map[T]bool, t T) map[T]bool {
		m[t] = true
		return m
	}, make(map[T]bool))
}

func SetDifference[K comparable](lhs map[K]bool, rhs map[K]bool) map[K]bool {
	var ret = make(map[K]bool)
	for c, v := range lhs {
		_, present := rhs[c]
		if v && !present {
			ret[c] = true
		}
	}
	return ret
}

func Frequencies[I iter.Seq[T], T comparable](iter I) map[T]int64 {
	return it.Fold(iter, func(m map[T]int64, t T) map[T]int64 {
		m[t]++
		return m
	}, make(map[T]int64))
}

func FlatMap[V, W any, S iter.Seq[W]](delegate func(func(V) bool), f func(V) S) iter.Seq[W] {
	return func(yield func(W) bool) {
		for innerValue := range delegate {
			for value := range f(innerValue) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

func FlatMap2[V, W, X, Y any, S iter.Seq2[X, Y]](delegate func(func(V, W) bool), f func(V, W) S) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for v, w := range delegate {
			for x, y := range f(v, w) {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

func Partition[T any](slice []T, n int, step int) iter.Seq[iter.Seq[T]] {
	ret := it.Exhausted[iter.Seq[T]]()
	for i := 0; i < len(slice); i += step {
		inner := slice[i:min(i+n, len(slice))]
		if len(inner) == n {
			ret = it.Chain(ret, it.Once(slices.Values(inner)))
		}
	}
	return ret
}
