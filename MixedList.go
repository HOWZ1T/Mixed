/*Copyright 2018 Dylan Randall. All rights reserved.
Use of this source code is governed by a Mozilla Public License 2.0
license that can be found in the LICENSE file. */
package MixedList

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// MixedList Errors:
//	- IndexOutOfBounds indicates that an given index has exceeded
//	  the boundaries of an array/list data type
var IndexOutOfBounds = errors.New("MixedList: IndexOutOfBounds: the given index exceeds the boundaries of the list")

// MixedList provides a list that is capable of storing values of varying types.
type MixedList []interface{}

// Has determines if the list has the item.
func (lst MixedList) Has(item interface{}) bool {
	for _, itm := range lst {
		if itm == item {
			return true
		}
	}

	return false
}

// Append add the item to end of the list.
func (lst *MixedList) Append(item interface{}) {
	*lst = append(*lst, item)
}

// Extend adds all items of another list to the list.
func (lst *MixedList) Extend(list MixedList) {
	*lst = append(*lst, list...)
}

// Insert adds the item at the given index into the list.
// If the given index exceeds the list a IndexOutOfBounds error is returned.
func (lst *MixedList) Insert(index int, item interface{}) error {
	if index < 0 || index > len(*lst) {
		return IndexOutOfBounds
	}

	*lst = append(*lst, 0)
	pval := *lst //pval -> pointer's value
	copy(pval[index+1:], pval[index:])
	pval[index] = item

	return nil
}

// Pop removes and returns the item at the given index from the list.
// If the given index exceeds the list an IndexOutOfBounds error is returned.
func (lst *MixedList) Pop(index int) (error, interface{}) {
	if index < 0 || index >= len(*lst) {
		return IndexOutOfBounds, nil
	}

	pval := *lst //pval -> pointer's value
	item := pval[index]

	*lst = append(pval[:index], pval[index+1:]...)
	return nil, item
}

// Clear removes all items from the list
func (lst *MixedList) Clear() {
	*lst = MixedList{}
}

// Index returns the index of the first matched item.
// If the item is not in the list a ValueError with an index of -1 is returned
func (lst MixedList) Index(item interface{}) (error, int) {
	for i := 0; i < len(lst); i++ {
		if lst[i] == item { return nil, i }
	}

	return fmt.Errorf("MixedList: ValueError: '%v' is not in the list", item), -1
}

// Size returns the length of the list
func (lst MixedList) Size() int {
	return len(lst)
}

// Reverse reverses the order of the items in the list
func (lst *MixedList) Reverse() {
	pval := *lst //pval -> pointer's value
	size := len(pval)

	for i := 0; i < size/2; i++ {
		j := size-i-1

		tmp := pval[j]
		pval[j] = pval[i]
		pval[i] = tmp
	}
}

// Copy returns a copy of the list
func (lst MixedList) Copy() MixedList {
	cpy := MixedList{}
	cpy.Extend(lst)
	return cpy
}

// Equals compares the list against another list for equality.
// Returns a bool, true denoting equal lists, and false denoting unequal lists
func (lst MixedList) Equals(list MixedList) bool {
	if len(lst) != len(list) { return false }

	for i := 0; i < len(lst); i++ {
		if lst[i] != list[i] { return false }
	}

	return true
}

// EqualsIgnoreCase similar too Equals however, when comparing strings it ignores case
func (lst MixedList) EqualsIgnoreCase(list MixedList) bool {
	if len(lst) != len(list) { return false }

	for i := 0; i < len(lst); i++ {
		if strings.ToLower(reflect.TypeOf(lst[i]).String()) == "string" && strings.ToLower(reflect.TypeOf(list[i]).String()) == "string" {
			sA := strings.ToLower(lst[i].(string))
			sB := strings.ToLower(list[i].(string))

			if sA != sB {
				return false
			}
		} else if lst[i] != list[i] {
			return false
		}
	}

	return true
}