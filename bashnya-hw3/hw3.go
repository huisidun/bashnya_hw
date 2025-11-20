package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	topIndex := len(s.data) - 1
	value := s.data[topIndex]
	s.data = s.data[:topIndex]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = nil 
}

func main() {
	var stack Stack

	fmt.Println("IsEmpty:", stack.IsEmpty()) 
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Size:", stack.Size()) 

	val, ok := stack.Pop()
	if ok {
		fmt.Println("Popped:", val) 
	}

	fmt.Println("Size after pop:", stack.Size()) 
	stack.Clear()
	fmt.Println("IsEmpty after clear:", stack.IsEmpty()) 
}