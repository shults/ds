package dlist

import (
	"errors"
)

var (
	ErrEmpty = errors.New("empty dlist")
)

type (
	// List represents double linked list implementation
	List[T any] struct {
		size int
		head *node[T]
		tail *node[T]
	}

	node[T any] struct {
		val  T
		prev *node[T]
		next *node[T]
	}
)

func New[T any]() *List[T] {
	return &List[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (list *List[T]) Empty() bool {
	return list.size == 0
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Prepend(val T) {
	var newNode = new(node[T])
	newNode.val = val

	switch list.size {
	case 0:
		list.head = newNode
		list.tail = newNode
	case 1:
		list.head = newNode
		list.head.next = list.tail
		list.tail.prev = list.head
	default:
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}

	list.size++
}

func (list *List[T]) Append(val T) {
	var newNode = new(node[T])
	newNode.val = val

	switch list.size {
	case 0:
		list.head = newNode
		list.tail = newNode
	case 1:
		list.tail = newNode
		list.head.next = list.tail
		list.tail.prev = list.head
	default:
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
}

func (list *List[T]) Head() (T, error) {
	var val T

	switch list.size {
	case 0:
		return val, ErrEmpty
	case 1:
		val = list.head.val
		list.head = nil
		list.tail = nil
	case 2:
		val = list.head.val
		list.head = list.tail
		list.head.next = nil
		list.head.prev = nil
	default:
		val = list.head.val
		list.head = list.head.next
		list.head.prev = nil
	}

	list.size--
	return val, nil
}

func (list *List[T]) Tail() (T, error) {
	var val T

	switch list.size {
	case 0:
		return val, ErrEmpty
	case 1:
		val = list.tail.val
		list.head = nil
		list.tail = nil
	case 2:
		val = list.tail.val
		list.tail = list.head
		list.tail.next = nil
		list.tail.prev = nil
	default:
		val = list.tail.val
		list.tail = list.head.prev
		list.tail.next = nil
	}

	list.size--
	return val, nil
}
