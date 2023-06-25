// https://exercism.org/tracks/go/exercises/list-ops

package main

func main() {

}

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	var result int = initial
	for _, value := range s {
		result = fn(result, value)
	}
	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	var result int = initial
	for _, value := range s.Reverse() {
		result = fn(value, result)
	}
	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	l := IntList{}
	for _, v := range s {
		if fn(v) {
			l = l.Append(IntList([]int{v}))
		}
	}
	return IntList(l)
}

func (s IntList) Length() int {
	var result int
	for range s {
		result++
	}
	return result
}

func (s IntList) Map(fn func(int) int) IntList {
	l := make([]int, len(s))
	for i, v := range s {
		l[i] = fn(v)
	}
	return IntList(l)
}

func (s IntList) Reverse() IntList {
	var result IntList = make([]int, 0)
	for _, value := range s {
		result = append(IntList{value}, result...)
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	l := make([]int, len(s)+len(lst))
	copy(l, s)
	for i, v := range lst {
		l[len(s)+i] = v
	}
	return IntList(l)
}

func (s IntList) Concat(lists []IntList) IntList {
	panic("Please implement the Concat function")
}
