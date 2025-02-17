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
func IsUnique[T comparable](arr []T) bool {
	seen := make(map[T]bool) // Map to track seen values
	for _, v := range arr {
		if seen[v] { // If the value is already in the map, it's a duplicate
			return false
		}
		seen[v] = true
	}
	return true
}
