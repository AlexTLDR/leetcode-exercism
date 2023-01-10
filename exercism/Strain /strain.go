package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var keep []int
	for _, v := range i {
		if filter(v) {
			keep = append(keep, v)
		}
	}
	return keep
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var discard []int
	for _, v := range i {
		if !filter(v) {
			discard = append(discard, v)
		}
	}
	return discard
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var keep [][]int
	for _, v := range l {
		if filter(v) {
			keep = append(keep, v)
		}
	}
	return keep
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var keep []string
	for _, v := range s {
		if filter(v) {
			keep = append(keep, v)
		}
	}
	return keep
}
