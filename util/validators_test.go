package util

import (
	"testing"
)

type Test struct {
	in  string
	out bool
}

var tests = []Test{
	{"email", false},
	{"email@", false},
	{"email@email.com", true},
}

func TestValidateEmail(t *testing.T) {
	for i, test := range tests {
		valid := ValidateEmail(test.in)
		if valid != test.out {
			t.Errorf("#%d: valid email(%s)=%v; want %v", i, test.in, valid, test.out)
		}
	}
}
