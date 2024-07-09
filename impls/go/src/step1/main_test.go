package main

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	input := `(+ 1 "hello" (myFun true false nil))`
	tokens := tokenize(input)

	expectedTokens := []Token{
		{
			Type:    LPAREN,
			Literal: "(",
		},
		{
			Type:    SYMBOL,
			Literal: "+",
		},
		{
			Type:    INT,
			Literal: "1",
		},
		{
			Type:    STRING,
			Literal: "hello",
		},
		{
			Type:    LPAREN,
			Literal: "(",
		},
		{
			Type:    SYMBOL,
			Literal: "myFun",
		},
		{
			Type:    TRUE,
			Literal: "true",
		},
		{
			Type:    FALSE,
			Literal: "false",
		},
		{
			Type:    NIL,
			Literal: "nil",
		},
		{
			Type:    RPAREN,
			Literal: ")",
		},
		{
			Type:    RPAREN,
			Literal: ")",
		},
		{
			Type:    EOF,
			Literal: "",
		},
	}

	testTokenLists(t, expectedTokens, tokens)
}

func testTokenLists(t *testing.T, expectedTokens, actual []Token) {
	t.Log(actual)
	if len(expectedTokens) != len(actual) {
		t.Fatalf("did not get the right number of tokens. expected=%d, got=%d", len(expectedTokens), len(actual))
	}

	for i, token := range actual {
		testToken(t, i, expectedTokens[i], token)
	}
}

func testToken(t *testing.T, tc int, expected, actual Token) {
	if actual.Type != expected.Type {
		t.Fatalf("token #%d, token type wrong. expected=%q, got=%q", tc, expected.Type, actual.Type)
	}

	if actual.Literal != expected.Literal {
		t.Fatalf("token #%d, literal wrong. expected=%q, got=%q", tc, expected.Type, actual.Literal)
	}
}

func TestReadStr(t *testing.T) {
	input := "(+ 1 (+ 2 (+ true false)))"
	ast := read_str(input)
	expectedOutputString := "(+ 1 (+ 2 (+ true false)))"

	if ast.String() != expectedOutputString {
		t.Fatalf("did not get a matching (string output) for the ast. expected=%s, got=%s", expectedOutputString, ast.String())
	}
	t.Log(ast)
}
