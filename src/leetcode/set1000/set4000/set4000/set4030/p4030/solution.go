package p4030

import "fmt"

func isPalindromic(s string) bool {
	t := make([]byte, 0, len(s)*8)
	for _, ch := range s {
		t = append(t, fmt.Sprintf("%0*b", 8, ch)...)
	}

	n := len(t)
	for i := range n / 2 {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true

}
