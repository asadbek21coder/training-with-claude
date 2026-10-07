package ds

import (
	imp "code/algorithms-ds/ds/ds_imp"
	"strings"
)

func IsValidParatheses(s string) bool {
	opening := "([{"
	closing := ")]}"
	stack := []rune{}
	for _, v := range s {
		if strings.Contains(opening, string(v)) {
			stack = append(stack, v)
		} else if strings.Contains(closing, string(v)) {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if string(last)+string(v) != "()" && string(last)+string(v) != "[]" && string(last)+string(v) != "{}" {
				return false
			}
		}

	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func IsValidParanthesesStack(s string) bool {
	stack := imp.NewStack()
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack.Push(v)
		case ')':
			if stack.IsEmpty() || stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.IsEmpty() || stack.Pop() != '[' {
				return false
			}

		case '}':
			if stack.IsEmpty() || stack.Pop() != '{' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

// MinStack - stack with O(1) GetMin operation
// Uses two stacks: main stack for values, min stack for tracking minimums at each step
type MinStack struct {
	main *IntStack
	min  *IntStack
}

type IntStack struct {
	items []int
}

func NewIntStack() *IntStack {
	return &IntStack{items: []int{}}
}

func (s *IntStack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *IntStack) Pop() int {
	if len(s.items) == 0 {
		return 0
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *IntStack) Peek() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}

func NewMinStack() *MinStack {
	return &MinStack{
		main: NewIntStack(),
		min:  NewIntStack(),
	}
}

func (s *MinStack) Push(val int) {
	s.main.Push(val)

	// Always push to min stack
	// If min stack is empty or val is smaller than current min, push val
	// Otherwise push current min again (to maintain alignment with main stack)
	if s.min.IsEmpty() || val < s.min.Peek() {
		s.min.Push(val)
	} else {
		s.min.Push(s.min.Peek())
	}
}

func (s *MinStack) Pop() int {
	s.min.Pop() // Always pop from min stack
	return s.main.Pop()
}

func (s *MinStack) Peek() int {
	return s.main.Peek()
}

func (s *MinStack) GetMin() int {
	return s.min.Peek()
}

// DailyTemperatures - monotonic stack approach, O(n) time
// For each day, find how many days until a warmer temperature
// Input:  [73, 74, 75, 71, 69, 72, 76, 73]
// Output: [ 1,  1,  4,  2,  1,  1,  0,  0]
func DailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	stack := []int{} // stores indices, not temperatures

	// Traverse from right to left (backwards)
	for i := len(temps) - 1; i >= 0; i-- {
		// Pop all indices where temperature <= current temperature
		// These can never be the next warmer day
		for len(stack) > 0 && temps[stack[len(stack)-1]] <= temps[i] {
			stack = stack[:len(stack)-1]
		}

		// If stack is not empty, top index is the next warmer day
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		} else {
			result[i] = 0
		}

		// Push current index onto stack
		stack = append(stack, i)
	}

	return result
}
