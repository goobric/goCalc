package calculator_test

import (
	calculator "myGoCalc"
	"testing"
)

/*
	 func TestAdd(t *testing.T) {
		t.Parallel()
		var want float64 = 4
		got := calculator.Add(2, 2)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
*/
func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 2, want: 3},
		{a: 5, b: 1, want: 6},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

/*
	 func TestSubtract(t *testing.T) {
		t.Parallel()
		var want float64 = 2
		got := calculator.Subtract(4, 2)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
*/
func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 5, b: 2, want: 3},
		{a: 1, b: 3, want: -2},
		{a: 5, b: 11, want: -6},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

/*
	 func TestMultiply(t *testing.T) {
		t.Parallel()
		var want float64 = 6
		got := calculator.Multiply(2, 3)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
*/
func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 0.5, want: 1},
		{a: -7, b: 2, want: -14},
		{a: 5.5, b: 3, want: 16.5},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: 1, b: -2, want: -0.5},
		{a: 12, b: 3, want: 4},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("failed, want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
