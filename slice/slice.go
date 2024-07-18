package slice

// NOTE: some of these methods assume that the slice is not empty

import (
	"math/rand"
	"time"
)

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	var _chunks = make([][]T, 0, (len(items)/chunkSize)+1)

	for chunkSize < len(items) {
		items, _chunks = items[chunkSize:], append(_chunks, items[0:chunkSize:chunkSize])
	}

	return append(_chunks, items)
}

func First[T any](input []T) T {
	return input[0]
}

func Rest[T any](input []T) []T {
	return input[1:]
}

func Last[T any](input []T) T {
	return input[len(input)-1]
}

func Reverse[T any](input []T) []T {
	reversed := []T{}

	for i := len(input) - 1; i >= 0; i-- {
		reversed = append(reversed, input[i])
	}

	return reversed
}

func RandomElement[T any](input []T) T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pick := r.Intn(len(input))
	return input[pick]
}

func IndexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T

	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}

	return result
}

func Map[T any, U any](list []T, f func(T) U) []U {
	output := []U{}

	for _, el := range list {
		output = append(output, f(el))
	}

	return output
}

func Filter[T any](list []T, f func(T) bool) []T {
	output := []T{}

	for _, el := range list {
		if f(el) {
			output = append(output, el)
		}
	}

	return output
}

func SliceAfter[T comparable](input []T, element T) []T {
	index := IndexOf(input, element)
	if index == -1 {
		return []T{}
	}
	return input[index+1:]
}

func Intersect[T comparable](input1 []T, input2 []T) []T {
	intersection := []T{}

	for _, element := range input1 {
		if IndexOf(input2, element) != -1 {
			intersection = append(intersection, element)
		}
	}

	return intersection
}

func Contains[T comparable](input []T, element T) bool {
	for _, el := range input {
		if el == element {
			return true
		}
	}
	return false
}
