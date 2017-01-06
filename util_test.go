package helloGOLib

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"Hello", "olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got, err := Reverse(c.in)
		if len(c.in) == 0 && err == nil {
			t.Errorf("Reverse(%q) == error, should be error", c.in)
		}
		if got != c.out {
			t.Errorf("Reverse(%q) == %q, should be %q", c.in, c.out, got)
		}
	}
}
