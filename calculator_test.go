package calculator_test

import (
	"math"
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
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
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
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
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
		{a: 1, b: 3, want: 0.33333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("failed, want no error for valid input, got %v", err)
		}
		// if tc.want != got {
		// 	t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		// }
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		input float64
		want  float64
	}
	testCases := []testCase{
		{input: 4, want: 2},
		{input: 2, want: 1.41421356237},
		{input: 25, want: 5},
		{input: 1.4, want: 1.2},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		if err != nil {
			t.Fatalf("Sqrt failed(%f): want no error for valid input, got %v", tc.input, err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt failed(%f): want %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-1)
	if err != nil {
		t.Error("Sqrt failed (-1): want error for valid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance

}

func TestAveMean(t *testing.T) {
	t.Parallel()
	_, err := calculator.AveMean(num)
	if err != nil {
		t.Error("Average Mean failed : want error for valid input, got nil")
	}
}
