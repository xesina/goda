package algorithms

import (
	"errors"
	"fmt"
)

func isBalancedBrackets(s string) bool {
	if s == "" {
		return true
	}
	const (
		openParentheses = "("
		openBracket     = "["
		openCurly       = "{"

		closeParentheses = ")"
		closeBracket     = "]"
		closeCurly       = "}"
	)

	type stack []string

	var stk stack

	pop := func() (string, error) {
		if len(stk) == 0 {
			return "", errors.New("stack is empty")
		}
		t := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		return t, nil
	}

	push := func(item string) {
		stk = append(stk, item)
	}

	isOpen := func(c string) bool {
		return c == openParentheses || c == openBracket || c == openCurly
	}

	balanced := true

	for _, c := range s {
		ch := fmt.Sprintf("%c", c)
		if isOpen(ch) {
			push(ch)
			continue
		}

		p, err := pop()
		if err != nil {
			balanced = false
		}

		switch ch {
		case closeParentheses:
			balanced = p == openParentheses

		case closeBracket:
			balanced = p == openBracket

		case closeCurly:
			balanced = p == openCurly

		default:
			balanced = false

		}

		if !balanced {
			break
		}

	}

	if _, err := pop(); err == nil {
		balanced = false
	}

	return balanced
}
