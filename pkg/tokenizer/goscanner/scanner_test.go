package goscanner_test

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/VixsTy/calculator/pkg/tokenizer"
	"github.com/VixsTy/calculator/pkg/tokenizer/goscanner"
	. "github.com/onsi/gomega"
)

func TestNominalScan(t *testing.T) {
	// Testcases
	tc := []struct {
		name     string
		input    string
		expected []struct {
			lit string
			tok token.Token
		}
		wantErr bool
	}{
		{
			name:  "without space",
			input: "1+1",
			expected: []struct {
				lit string
				tok token.Token
			}{
				{"1", token.INT},
				{"+", token.ADD},
				{"1", token.INT},
			},
			wantErr: false,
		},
		{
			name:  "with space",
			input: "1 + 1",
			expected: []struct {
				lit string
				tok token.Token
			}{
				{"1", token.INT},
				{"+", token.ADD},
				{"1", token.INT},
			},
			wantErr: false,
		},
		{
			name:  "with parenthesis",
			input: "(125 * 7) / 36",
			expected: []struct {
				lit string
				tok token.Token
			}{
				{"(", token.LPAREN},
				{"125", token.INT},
				{"*", token.MUL},
				{"7", token.INT},
				{")", token.RPAREN},
				{"/", token.QUO},
				{"36", token.INT},
			},
			wantErr: false,
		},
		{
			name:  "with float",
			input: "(125.68 * 7.25) / 36",
			expected: []struct {
				lit string
				tok token.Token
			}{
				{"(", token.LPAREN},
				{"125.68", token.FLOAT},
				{"*", token.MUL},
				{"7.25", token.FLOAT},
				{")", token.RPAREN},
				{"/", token.QUO},
				{"36", token.INT},
			},
			wantErr: false,
		},
		{
			name:  "with EOF",
			input: "",
			expected: []struct {
				lit string
				tok token.Token
			}{
				{"", token.EOF},
			},
			wantErr: false,
		},
		{
			name:  "with EOF",
			input: "HELLO",
			expected: []struct {
				lit string
				tok token.Token
			}{
				{"HELLO", token.STRING},
			},
			wantErr: true,
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
				lit, tok, err := parser.Scan()
				if testCase.wantErr {
					g.Expect(err).To(Equal(tokenizer.UnrecognizedToken))
				} else {
					g.Expect(err).To(BeNil())
					g.Expect(tok).To(Equal(v.tok), fmt.Sprintf("Expected %s %s be equal %s %s", lit, tok.String(), v.lit, v.tok.String()))
					g.Expect(lit).To(Equal(v.lit))
				}
			}
		})
	}
}
