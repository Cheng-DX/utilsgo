package utilsgo

import (
	"fmt"
)

func At[T any](collection []T, index int) (T, bool) {
	if index < 0 || index >= len(collection) {
		var r T
		return r, false
	}

	return collection[index], true
}

func Entries[T any](collection []T) []Tuple_2[int, T] {
	return Map(collection, func(item T, index int) Tuple_2[int, T] {
		return Tuple_2[int, T]{
			One: index,
			Two: item,
		}
	})
}

func Filter[T any](collection []T, condition func(item T, index int) bool) []T {
	r := []T{}

	for index, item := range collection {
		if condition(item, index) {
			r = append(r, item)
		}
	}
	return r
}

func Find[T any](collection []T, tester func(item T) bool) (T, bool) {
	for _, item := range collection {
		if tester(item) {
			return item, true
		}
	}
	var r T
	return r, false
}

func FindIndexOf[T any](collection []T, tester func(item T) bool) (int, bool) {
	for index, item := range collection {
		if tester(item) {
			return index, true
		}
	}
	return -1, false
}

func FindLast[T any](collection []T, tester func(item T) bool) (T, bool) {
	len := len(collection)
	for i := len - 1; i >= len; i-- {
		if tester(collection[i]) {
			return collection[i], true
		}
	}
	var r T
	return r, false
}

func FindLastIndexOf[T any](collection []T, tester func(item T) bool) (int, bool) {
	len := len(collection)
	for i := len - 1; i >= len; i-- {
		if tester(collection[i]) {
			return i, true
		}
	}
	return -1, false
}

func Map[T any, R any](collection []T, mapper func(item T, index int) R) []R {
	r := []R{}
	for index, item := range collection {
		r = append(r, mapper(item, index))
	}
	return r
}

func Every[T any](collection []T, condition func(item T) bool) bool {
	for _, item := range collection {
		if !condition(item) {
			return false
		}
	}

	return true
}

func Some[T any](collection []T, condition func(item T) bool) bool {
	for _, item := range collection {
		if condition(item) {
			return true
		}
	}

	return false
}

func ForEach[T any](collection []T, handler func(item T, index int)) {
	for i, item := range collection {
		handler(item, i)
	}
}

func Slice[T any](collection []T, startIndex int, endIndex int) []T {
	length := len(collection)
	if startIndex < -length || startIndex > length {
		return []T{}
	}
	if endIndex < 0 {
		endIndex += length
	}
	if startIndex < 0 {
		startIndex += length
	}
	if endIndex > len(collection) {
		endIndex = len(collection)
	}
	return collection[startIndex:endIndex]
}

func Join[T any](collection []T, seperator string) string {
	result := ""
	for index, item := range collection {
		result += fmt.Sprintf("%v", item)
		if index < len(collection)-1 {
			result += seperator
		}
	}
	return result
}

func Includes[T comparable](collection []T, target T) bool {
	for _, item := range collection {
		if item == target {
			return true
		}
	}
	return false
}

func Pop[T any](collection []T) ([]T, T) {
	length := len(collection)
	if length == 0 {
		var t T
		return collection, t
	}
	return collection[:length-1], collection[length-1]
}

func Push[T any](collection []T, item T) []T {
	return append(collection, item)
}

func Shift[T any](collection []T) ([]T, T) {
	length := len(collection)
	if length == 0 {
		var t T
		return collection, t
	}
	return collection[1:], collection[0]
}

func Reverse[T any](collection []T) []T {
	length := len(collection)
	if length == 0 || length == 1 {
		return collection
	}
	result := []T{}

	for i := length - 1; i >= 0; i-- {
		result = append(result, collection[i])
	}
	return result
}
