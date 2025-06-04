package lexer

import (
	"testing"

	"github.com/ienjir/Switzerlang/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	testCases := []struct {
		expectedType     token.TokenType
		exptectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, testCase := range testCases {
		token := 	lexer.NextToken()

		if token.Type != testCase.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected: %q, got: %q", i, &testCase.expectedType, token.Type)
		}

		if token.exptectedLiteral != testCase.exptectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected: %q, got: %q", i, testCase.exptectedLiteral, token.Literal)
		}
	}
}
