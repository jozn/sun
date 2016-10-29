package ds

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraylist implements the array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
//package arraylist

import (
	"fmt"
	//"github.com/emirpasic/gods/lists"
	//"github.com/emirpasic/gods/utils"
	"sort"
	"strings"
)

//func assertListImplementation() {
//	var _ lists.List = (*List)(nil)
//}

// IntList holds the elements in a slice
type IntList struct {
	//elements []interface{}
	Elements []int
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new empty list
func New() *IntList {
	return &IntList{}
}

//func New() IntList {
//    return IntList{}
//}

// Add appends a value at the end of the list
func (list *IntList) Add(values ...int) {
	list.growBy(len(values))
	for _, value := range values {
		list.Elements[list.size] = value
		list.size++
	}
}

func (list *IntList) AddAndSort(values ...int) {
	list.Add(values...)
	list.SortDesc()
}

func (list *IntList) RemoveAndSort(value int) {
	list.RemoveIndex(value)
	for {
		n := sort.Search(list.size, func(i int) bool { return list.Elements[i] <= value })
		if n >= list.size || n < 0 {
			break
		}
		v := list.Elements[n]
		if v == value {
			list.RemoveIndex(n)
		} else {
			break
		}
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *IntList) Get(index int) (int, bool) {

	if !list.withinRange(index) {
		return 0, false
	}

	return list.Elements[index], true
}

// Remove removes one or more elements from the list with the supplied indices.
//ME: no need resort
func (list *IntList) RemoveIndex(index int) {

	if !list.withinRange(index) {
		return
	}

	list.Elements[index] = 0                                      // cleanup reference
	copy(list.Elements[index:], list.Elements[index+1:list.size]) // shift to the left by one (slow operation, need ways to optimize this)
	list.size--

	list.shrink()
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *IntList) Contains(values ...int) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range list.Elements {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

//must be sorted asc
func (list *IntList) BinaryContains(value int) bool {
	//n:= sort.SearchInts(list.Aelements, value)
	//list.SortAsc()
	//n:=sort.SearchInts(list.Elements, value)
	n := sort.Search(list.size, func(i int) bool { return list.Elements[i] <= value })
	//print("Search : ",n," ", value ," ", list.size, " ", list.Elements , " ", len(list.Elements), "\n")
	//fmt.Println(list.Elements)
	//list.SortDesc()
	if n >= list.size || n < 0 {
		return false
	}
	if list.Elements[n] == value {
		return true
	}
	//copy()
	return false
}

// Values returns all elements in the list.
func (list *IntList) Values() []int {
	newElements := make([]int, list.size, list.size)
	copy(newElements, list.Elements[:list.size])
	return newElements
}

// Empty returns true if list does not contain any elements.
func (list *IntList) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *IntList) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *IntList) Clear() {
	list.size = 0
	list.Elements = []int{}
}

// Sort sorts values (in-place) using.
/*func (list *IntList) Sort(comparator Comparator) {
	if len(list.elements) < 2 {
		return
	}
	Sort(list.elements[:list.size], comparator)
}*/

//func (list *IntList) Sort() {
//    sort.Sort(list)
//}

// Swap swaps the two values at the specified positions.
func (list *IntList) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.Elements[i], list.Elements[j] = list.Elements[j], list.Elements[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *IntList) Insert(index int, values ...int) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	// Shift old to right
	for i := list.size - 1; i >= index+l; i-- {
		list.Elements[i] = list.Elements[i-l]
	}
	// Insert new
	for i, value := range values {
		list.Elements[index+i] = value
	}
}

// String returns a string representation of container
func (list *IntList) String() string {
	str := "ArrayIntList\n"
	values := []string{}
	for _, value := range list.Elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *IntList) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *IntList) resize(cap int) {
	newElements := make([]int, cap, cap)
	copy(newElements, list.Elements)
	list.Elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *IntList) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.Elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *IntList) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.Elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

//By me For sorting
func (p *IntList) SortAsc() { sort.Sort(p) }

func (p *IntList) SortDesc() {
	sort.Sort(sort.Reverse(p))
}

func (p *IntList) TrySortDesc() {
	if !sort.IsSorted(sort.Reverse(p)) {
		p.SortDesc()
	}
}

func (p *IntList) Len() int           { return p.size /*len(p.elements)*/ }
func (p *IntList) Less(i, j int) bool { return p.Elements[i] < p.Elements[j] }

//func (p *IntList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
