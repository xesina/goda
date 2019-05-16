package algorithms

// Checking for balanced brackets using stack data structure
// TODO: refactor this function with stack already written in stack package

import "testing"

func TestIsBalancedBrackets(t *testing.T) {
	cases := []struct {
		expr string
		want bool
	}{
		{expr: "()", want: true},
		{expr: "((", want: false},
		{expr: "{", want: false},
		{expr: "(({{[]}}))", want: true},
		{expr: "(({[})]", want: false},
		{expr: "", want: true},
		{expr: "([]", want: false},
		{expr: "{{)}", want: false},
		{expr: "{{{{{{{{{{{{{{{{{", want: false},
		{expr: "((()))", want: true},
	}

	for _, c := range cases {
		if b := isBalancedBrackets(c.expr); b != c.want {
			t.Errorf(`isBalancedBrackets("%s") = %v, want %v`, c.expr, b, c.want)
		}
	}
}
