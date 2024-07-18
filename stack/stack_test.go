package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	if len(stack.items) != 1 || stack.items[0] != 1 {
		t.Errorf("Push failed. Expected [1], got %v", stack.items)
	}

	stack.Push(2)
	if len(stack.items) != 2 || stack.items[1] != 2 {
		t.Errorf("Push failed. Expected [1 2], got %v", stack.items)
	}

	// Add more test cases here if needed

}
func TestPop(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	popped := stack.Pop()
	if popped != 3 {
		t.Errorf("Pop failed. Expected 3, got %v", popped)
	}

	popped = stack.Pop()
	if popped != 2 {
		t.Errorf("Pop failed. Expected 2, got %v", popped)
	}

	popped = stack.Pop()
	if popped != 1 {
		t.Errorf("Pop failed. Expected 1, got %v", popped)
	}

	popped = stack.Pop()
	if popped != 0 {
		t.Errorf("Pop failed. Expected 0, got %v", popped)
	}
}
func TestPopAll(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	items := stack.PopAll()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(items, expected) {
		t.Errorf("PopAll failed. Expected %v, got %v", expected, items)
	}

	if len(stack.items) != 0 {
		t.Errorf("PopAll failed. Expected stack to be empty, but it still contains %v", stack.items)
	}
}
func TestPopAllInReverse(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	items := stack.PopAllInReverse()
	expected := []int{3, 2, 1}

	if !reflect.DeepEqual(items, expected) {
		t.Errorf("PopAllInReverse failed. Expected %v, got %v", expected, items)
	}

	if len(stack.items) != 0 {
		t.Errorf("PopAllInReverse failed. Expected stack to be empty, but it still contains %v", stack.items)
	}
}
