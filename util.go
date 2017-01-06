package helloGOLib

import "errors"

// Reverse returns the string in reverse
func Reverse(s string) (string, error) {
	if len(s) == 0 {
		return s, errors.New("String is empty")
	}

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r), nil
}
