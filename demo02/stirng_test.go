package demo02

import "testing"

func TestString(t *testing.T) {
	s := "磨"
	t.Log(len(s))
	rs := []rune(s)
	t.Logf("type: %[1]T; value %[1]x", rs)
}
