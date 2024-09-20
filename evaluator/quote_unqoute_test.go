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
