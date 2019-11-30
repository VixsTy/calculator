package shuntingyard_test

import (
	"testing"

	shuntingyard "github.com/VixsTy/calculator/pkg/calculator/shunting-yard"
	. "github.com/onsi/gomega"
)

func TestNominalCalc(t *testing.T) {
	// Testcases
	tc := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "without addition",
			input:    "1+2",
			expected: "3",
			wantErr:  false,
		},
		{
			name:     "with soustraction",
			input:    "2 + 4",
			expected: "6",
			wantErr:  false,
		},
		{
			name:     "with parenthesis",
			input:    "(125 * 7) / 36",
			expected: "24.305555555555557",
			wantErr:  false,
		},
		{
			name:     "with float",
			input:    "(125.68 * 7.25) / 36",
			expected: "25.310555555555556",
			wantErr:  false,
		},
		{
			name:     "with modulo",
			input:    "47 % 6",
			expected: "5",
			wantErr:  false,
		},
		{
			name:     "with pow",
			input:    "2^2",
			expected: "4",
			wantErr:  false,
		},
		{
			name:     "with multiple operation",
			input:    "3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3",
			expected: "3.0001220703125",
			wantErr:  false,
		},
	}

	// Run as subtests
	for _, tt := range tc {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			// t.Parallel()

			g := NewGomegaWithT(t)

			calc := shuntingyard.NewShuntingYard()

			result := calc.Calc(testCase.input)
			g.Expect(result).To(Equal(testCase.expected))
		})
	}
}
