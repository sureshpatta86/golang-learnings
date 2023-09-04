package main

import "testing"

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		if Add(a, b) != a+b {
			t.Errorf("Add(%d, %d) != %d", a, b, a+b)
		}
	})
}
