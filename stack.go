package stack

type Stack[T any] []T

func New[T any]() Stack[T] {
	return Stack[T]{}
}

// Push an element onto the stack in O(1) time.
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Peek at the top element of the stack in O(1) time.
func (s *Stack[T]) Top() T {
	return (*s)[len(*s)-1]
}

// Pop an element from the stack in O(1) time.
func (s *Stack[T]) Pop() {
	*s = (*s)[:len(*s)-1]
}

// Pop multiple elements from the stack in O(1) time.
func (s *Stack[T]) PopN(n int) {
	*s = (*s)[:len(*s)-n]
}
