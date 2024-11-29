package utils

// Generic Stack implementation using a slice
type Stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top element of the stack. Returns zero value
// of T if stack is empty.
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	index := len(s.data) - 1
	element := s.data[index]
	s.data = s.data[:index]
	return element
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) ToSlice() []T {
	return s.data
}

/*
func main() {
	// Create a stack of integers
	intStack := new(Stack[int])
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	for !intStack.IsEmpty() {
		fmt.Println(intStack.Pop())
	}

	// Create a stack of strings
	stringStack := new(Stack[string])
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Pop())
}
*/
