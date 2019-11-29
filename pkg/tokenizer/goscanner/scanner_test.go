package goscanner_test

import (
	"io"
	"testing"

	"github.com/VixsTy/calculator/pkg/tokenizer/goscanner"
	. "github.com/onsi/gomega"
)

func TestNominalScan(t *testing.T) {
	// Testcases
	tc := []struct {
		name     string
		input    string
		expected []string
		wantErr  bool
	}{
		{
			name:     "without space",
			input:    "1+1",
			expected: []string{"1", "+", "1"},
			wantErr:  false,
		},
		{
			name:     "with space",
			input:    "1 + 1",
			expected: []string{"1", "+", "1"},
			wantErr:  false,
		},
		{
			name:     "with parenthesis",
			input:    "(125 * 7) / 36",
			expected: []string{"(", "125", "*", "7", ")", "/", "36"},
			wantErr:  false,
		},
		{
			name:     "with EOF",
			input:    "",
			expected: []string{"EOF"},
			wantErr:  true,
		},
	}

	// Run as subtests
	for _, tt := range tc {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			// t.Parallel()

			g := NewGomegaWithT(t)

			parser := &goscanner.GoScanner{}
			parser.Init(testCase.input)

			// assert results expectations
			for _, v := range testCase.expected {
				token, err := parser.Scan()
				if testCase.wantErr {
					g.Expect(err).To(Equal(io.EOF))
				} else {
					g.Expect(err).To(BeNil())
					g.Expect(token).To(Equal(v))
				}
			}
		})
	}
}
