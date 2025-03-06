package stack

import "errors"

var (
	ErrEmpty = errors.New("empty stack")
)

type Stack[T any] struct {
	size int
	next *node[T]
}

type node[T any] struct {
	val  T
	next *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		size: 0,
		next: nil,
	}
}

func (stack *Stack[T]) Size() int {
	return stack.size
}

func (stack *Stack[T]) Empty() bool {
	return stack.size == 0
}

func (stack *Stack[T]) Push(val T) {
	var newNode = new(node[T])
	newNode.val = val
	newNode.next = stack.next
	stack.next = newNode
	stack.size++
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.Empty() {
		var zero T
		return zero, ErrEmpty
	}

	stack.size--
	var val = stack.next.val
	stack.next = stack.next.next

	return val, nil
}
