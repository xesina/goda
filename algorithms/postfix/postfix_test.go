package postfix

import "testing"

func TestInfixToPostfix(t *testing.T) {
	cases := []struct {
		expr string
		want string
	}{
		{
			expr: "1+2",
			want: "12+",
		},
		{
			expr: "A+B*C-D*E",
			want: "ABC*+DE*-",
		},
		{
			expr: "a+b-c",
			want: "ab+c-",
		},
		{
			expr: "a+b-c/d*e*f/g-h",
			want: "ab+cd/e*f*g/-h-",
		},
	}

	for _, c := range cases {
		w := InfixToPostfix(c.expr)
		if w != c.want {
			t.Errorf(`InfixToPostfix("%s") = "%s", want "%s"`, c.expr, w, c.want)
		}
	}
}
