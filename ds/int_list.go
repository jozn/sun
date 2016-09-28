package ds

import "sort"

type IntList_______DEppppp []int

func NewIntList() IntList_______DEppppp {
	return IntList_______DEppppp{}
}

func (l IntList_______DEppppp) Add(vals ...int) {
	for _, v := range vals {
		l = append(l, v)
	}
}

func (l IntList_______DEppppp) Remove(vals ...int) {
	for _, v := range vals {
		l = append(l, v)
	}
}

func (l IntList_______DEppppp) Contains(vals ...int) {
	for _, v := range vals {
		l = append(l, v)
	}
}

func (p IntList_______DEppppp) Len() int           { return len(p) }
func (p IntList_______DEppppp) Less(i, j int) bool { return p[i] < p[j] }
func (p IntList_______DEppppp) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p IntList_______DEppppp) SortAsc()  { sort.Sort(p) }
func (p IntList_______DEppppp) SortDesc() { sort.Sort(sort.Reverse(p)) }
