package evaluator

import (
	"fmt"
	"testing"

	"github.com/ath0m/monkey/object"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`quote(5)`,
			`5`,
		},
		{
			`quote(5 + 8)`,
			`(5 + 8)`,
		},
		{
			`quote(foobar)`,
			`foobar`,
		},
		{
			`quote(foobar + barfoo)`,
			`(foobar + barfoo)`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			evaluated := testEval(tt.input)
			quote, ok := evaluated.(*object.Quote)
			if !ok {
				t.Fatalf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
			}
			if quote.Node == nil {
				t.Fatal("quote.Node is nil")
			}
			if quote.Node.String() != tt.expected {
				t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
			}
		})
	}
}

func TestQuoteUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`quote(unquote(4))`,
			`4`,
		},
		{
			`quote(unquote(4 + 4))`,
			`8`,
		},
		{
			`quote(8 + unquote(4 + 4))`,
			`(8 + 8)`,
		},
		{
			`quote(unquote(4 + 4) + 8)`,
			`(8 + 8)`,
		},
		{
			`let foobar = 8;
			quote(foobar)`,
			`foobar`,
		},
		{
			`let foobar = 8;
			quote(unquote(foobar))`,
			`8`,
		},
		{
			`quote(unquote(true))`,
			`true`,
		},
		{
			`quote(unquote(true == false))`,
			`false`,
		},
		{
			`quote(unquote(quote(4 + 4)))`,
			`(4 + 4)`,
		},
		{
			`let quotedInfixExpression = quote(4 + 4);
			quote(unquote(4 + 4) + unquote(quotedInfixExpression))`,
			`(8 + (4 + 4))`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			evaluated := testEval(tt.input)
			quote, ok := evaluated.(*object.Quote)
			if !ok {
				t.Fatalf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
			}
			if quote.Node == nil {
				t.Fatal("quote.Node is nil")
			}
			if quote.Node.String() != tt.expected {
				t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
			}
		})
	}
}
