package stack

func NewStringStack() *Stack[string] {
	return &Stack[string]{
		items: []string{},
	}
}
