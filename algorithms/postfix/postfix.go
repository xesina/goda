package postfix

import (
	"fmt"
)

var (
	operators = map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
)

type stack struct {
	data []string
}

func (s *stack) push(item string) {
	s.data = append(s.data, item)
}

func (s *stack) pop() string {
	if len(s.data) == 0 {
		return ""
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]

	return top
}

func (s *stack) top() string {
	return s.data[len(s.data)-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

// InfixToPostfix converts a infix expression to a postfix expression
func InfixToPostfix(expr string) string {
	s := &stack{}
	postfix := ""
	for _, c := range expr {
		ch := fmt.Sprintf("%c", c)
		_, isOperator := operators[ch]
		if !isOperator {
			postfix += ch
			continue
		}

		for !s.IsEmpty() && operators[s.top()] >= operators[ch] {
			postfix += s.pop()
		}

		s.push(ch)
	}

	for !s.IsEmpty() {
		postfix += s.pop()
	}

	return postfix
}
