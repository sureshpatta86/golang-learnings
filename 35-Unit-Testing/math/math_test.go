package math

import (
	"os"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {

	// if Abs(-1) < 0 {
	// 	t.Error("Negative Value found in the Abs method with input parameter", -1)
	// }

	// if Abs(0) < 0 {
	// 	t.Error("Negative Value found in the Abs method with input parameter", 0)
	// }

	// if Abs(1) < 0 {
	// 	t.Error("Negative Value found in the Abs method with input parameter", 1)
	// }
}

func TestAbsSub(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if Abs(1) < 0 {
			t.Error("Negative Value found in the Abs method with input parameter", 1)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		if Abs(0) < 0 {
			t.Error("Negative Value found in the Abs method with input parameter", 0)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		if Abs(-1) < 0 {
			t.Error("Negative Value found in the Abs method with input parameter", -1)
		}
	})

}

func TestSkip(t *testing.T) {

	if len(os.Getenv("GOPATH")) == 0 {
		t.Skip("Skipping the test case due to GOPATH is not set")
	}
	t.Log("Tested with GOPATH", os.Getenv("GOPATH"))

}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Positive", args: args{a: 1, b: 2}, want: 3},
		{name: "Negative", args: args{a: -1, b: -2}, want: -3},
		{name: "Zero", args: args{a: 0, b: 0}, want: 0},
		{name: "Mixed", args: args{a: -1, b: 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombineString(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Positive", args: args{a: "Hello", b: "World"}, want: "HelloWorld"},
		{name: "Negative", args: args{a: "Hello", b: " World"}, want: "Hello World"},
		{name: "Empty", args: args{a: "", b: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineString(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CombineString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanUp(t *testing.T) {
	t.Cleanup(func() { //it will be called last and works like defer
		t.Log("Cleanup called")
	})
	t.Log("Test started") //it will be called first
}

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestParallelThree(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
