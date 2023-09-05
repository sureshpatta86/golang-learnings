package main

import "testing"

type testCase struct {
	a int
	b int
	sum int
}

func TestAdd(t *testing.T) {

	tc := testCase{2, 2, 2+2}
	expected := tc.sum
	actual := Add(tc.a, tc.b)

	if actual != expected {
		t.Errorf("Add(%d, %d): expected %d, actual %d", tc.a, tc.b, expected, actual)
	}	
}
