package list

import "github.com/joeshaw/gengen/generic"

type List struct {
	data generic.T
	next *List
}

func (l *List) Prepend(d generic.T) *List {
	n := &List{
		data: d,
		next: l,
	}

	return n
}

func (l *List) Contains(d generic.T) bool {
	if l == nil {
		return false
	}

	for i := l; i != nil; i = i.next {
		if i.data == d {
			return true
		}
	}

	return false
}

func (l *List) Data() generic.T {
	if l == nil {
		// Return the zero value for generic.T, whatever type it ends
		// up becoming
		var zero generic.T
		return zero
	}

	return l.data
}
