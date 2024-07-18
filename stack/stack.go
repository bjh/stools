// https://rksurwase.medium.com/slice-based-stack-implementation-in-golang-8140603a1dc2
package stack

import "github.com/bjh/stools/slice"

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

func (stack *Stack[T]) Pop() T {
	var noValue T
	n := len(stack.items)

	if n <= 0 {
		return noValue
	}

	lastItem := slice.Last(stack.items)
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

func (stack *Stack[T]) PopAll() []T {
	items := stack.items
	stack.items = []T{}

	return items
}

func (stack *Stack[T]) PopAllInReverse() []T {
	items := stack.items
	stack.items = []T{}

	return slice.Reverse(items)
}
