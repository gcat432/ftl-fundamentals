package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		inputs      []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Sum of some positive numbers",
			inputs: []float64{2, 3, 4}, want: 9},
		{name: "Sum of some negative numbers",
			inputs: []float64{-2, -3, -4}, want: -9},
		{name: "Sum of some fractional numbers",
			inputs: []float64{5.5, -3.2, 9.33, 512.2}, want: 523.83},
		{name: "Sum of one operand",
			inputs: []float64{2}, want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Add(tc.inputs...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Add(%v): unexpected error status: %v",
				tc.inputs, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Add(%v): %s: want %f, got %f",
				tc.inputs, tc.name, tc.want, got)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	type testCase struct {
		name   string
		inputs []float64
		want   float64
	}

	tc := testCase{name: "Sum of some positive numbers",
		inputs: []float64{2, 3, 4}}

	for i := 0; i < b.N; i++ {
		_, _ = calculator.Add(tc.inputs...)
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		inputs      []float64
		want        float64
		errExpected bool
	}

	for i := 0; i < 100; i++ {
		tc := testCase{}

		loops := rand.Intn(10) + 1 // At least one operand
		if loops < 2 {
			tc.errExpected = true
		}

		for j := 0; j < loops; j++ {
			nb := rand.Float64()
			tc.inputs = append(tc.inputs, nb)
			tc.want += nb
		}

		got, err := calculator.Add(tc.inputs...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Add(%v): unexpected error status: %v",
				tc.inputs, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Add(%v): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		inputs      []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Difference of some positive numbers",
			inputs: []float64{2, 3, 4}, want: -5},
		{name: "Difference of some negative numbers",
			inputs: []float64{-2, -3, -4}, want: 5},
		{name: "Difference of some fractional numbers",
			inputs: []float64{5.5, -3.2, 9.33, 512.2}, want: -512.83},
		{name: "Difference of one operand",
			inputs: []float64{2}, want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Substract(tc.inputs...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Substract(%v): unexpected error status: %v",
				tc.inputs, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Substract(%v): %s: want %f, got %f",
				tc.inputs, tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		inputs      []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Product of some positive numbers",
			inputs: []float64{2, 3, 4}, want: 24},
		{name: "Product of some negative numbers",
			inputs: []float64{-2, -3, -4}, want: -24},
		{name: "Product of some fractional numbers",
			inputs: []float64{5.5, -3.2, 9.33, 512.2}, want: -84107.338},
		{name: "Product of one operand",
			inputs: []float64{2}, want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Multiply(tc.inputs...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Multiply(%v): unexpected error status: %v",
				tc.inputs, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Multiply(%v): %s: want %f, got %f",
				tc.inputs, tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		inputs      []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Quotient of some positive numbers",
			inputs: []float64{6, 3, 2}, want: 1},
		{name: "Quotient of some negative numbers",
			inputs: []float64{-6, -3, -1}, want: -2},
		{name: "Product of some fractional numbers",
			inputs: []float64{6.5, -2.5, 1.3}, want: -2},
		{name: "Quotient of some positive numbers by zero",
			inputs: []float64{3, 2, 0, 1}, want: 999, errExpected: true},
		{name: "Quotient of one operand",
			inputs: []float64{3}, want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.inputs...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%v): unexpected error status: %v",
				tc.inputs, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%v): %s: want %f, got %f",
				tc.inputs, tc.name, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		inputs      []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Square root of a positive number",
			inputs: []float64{64}, want: 8},
		{name: "Square root of a fractional number",
			inputs: []float64{0.64}, want: 0.8},
		{name: "Square root of a negative number",
			inputs: []float64{-64}, want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.inputs[0])
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Sqrt(%f): unexpected error status: %v",
				tc.inputs[0], errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Sqrt(%f): %s: want %f, got %f",
				tc.inputs[0], tc.name, tc.want, got)
		}
	}
}

func TestCalculate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		input       string
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Sum of numbers in a string",
			input: "1 + 1.5", want: 2.5},
		{name: "Difference of numbers in a string",
			input: "100-0.1", want: 99.9},
		{name: "Product of numbers in a string",
			input: "2*2", want: 4},
		{name: "Quotient of numbers in a string",
			input: "18  /  6", want: 3},
		{name: "Quotient of a number by zero in a string",
			input: "18 / 0", want: 999, errExpected: true},
		{name: "Unknown operator in a string",
			input: "18 @ 0", want: 999, errExpected: true},
		{name: "Wrong formatting in the first operand of a string",
			input: "abc + 10", want: 999, errExpected: true},
		{name: "Wrong formatting in the second operand of a string",
			input: "-10 - abc", want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Calculate(tc.input)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Calculate(%s): unexpected error status: %v",
				tc.input, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Calculate(%s): %s: want %f, got %f",
				tc.input, tc.name, tc.want, got)
		}
	}
}
