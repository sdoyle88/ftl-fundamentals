package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 2, b: 3, want: 5},
		{a: 5, b: 5, want: 10},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 2, b: 3, want: -1},
		{a: 5, b: 3, want: 2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.b, tc.a)
		if tc.want != got {
			t.Errorf("Subtract(%f,%f): want %f, got %f", tc.b, tc.a, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 2, b: 3, want: 6},
		{a: 5, b: 3, want: 15},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 1, errExpected: false},
		{a: 2, b: 0, want: 0, errExpected: true},
		{a: 6, b: 3, want: 2, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide (%f,%f): unexpected error status: %v", tc.a, tc.b, errReceived)
		} else if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
