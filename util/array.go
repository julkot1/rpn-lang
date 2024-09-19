package util

import "reflect"

func ReverseSlice(slice interface{}) {

	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("reverseSlice: not a slice")
	}

	length := s.Len()

	swap := reflect.Swapper(slice)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
