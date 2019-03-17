package gstuff

import (
	"reflect"
)

// Array ..
type Array struct {
}

// MaxInt ..
func (Array) MaxInt(arr []int) (max int) {
	for i := range arr {
		if i > max {
			max = i
		}
	}
	return
}

// MinInt ..
func (Array) MinInt(arr []int) (min int) {
	for i := range arr {
		if i < min {
			min = i
		}
	}
	return
}

// Index ..
func (Array) Index(arr []interface{}, item interface{}) (index int) {
	s := reflect.ValueOf(arr)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(item, s.Index(i).Interface()) == true {
			index = i
			break
		}
	}
	return
}
