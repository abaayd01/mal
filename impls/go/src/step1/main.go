package main

// import "runtime"

// "github.com/chzyer/readline"

// func main() {
// 	rl, err := readline.New("user> ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rl.Close()
//
// 	for {
// 		line, err := rl.Readline()
// 		if err != nil {
// 			break
// 		}
//
// 		result := rep(line)
// 		println(result)
// 	}
// }

func rep(in string) string {
	return _print(eval(read(in)))
}

func read(in string) string {
	return in
}

func eval(in string) string {
	return in
}

func _print(in string) string {
	return in
}

func read_str(in string) MalType {
	return nil
}

func tokenize(in string) []Token {
	tokens := []Token{}
	l := NewLexer(in)

	for {
		t := l.NextToken()
		if t.Type == EOF {
			tokens = append(tokens, t)
			break
		}
		tokens = append(tokens, t)
	}

	return tokens
}

func read_form(reader *Reader) MalType {
	return nil
}

func read_atom(reader *Reader) MalType {
	return nil
}

type Reader struct {
	curPosition int
	tokens      []Token
}

func (r *Reader) peek() Token {
	if r.curPosition > len(r.tokens) {
		return Token{
			Type:    EOF,
			Literal: "",
		}
	}

	return r.tokens[r.curPosition]
}

func (r *Reader) next() Token {
	if r.curPosition > len(r.tokens) {
		return Token{
			Type:    EOF,
			Literal: "",
		}
	}

	token := r.tokens[r.curPosition]
	r.curPosition++
	return token
}

type MalType interface{}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // the original string from the source that generated this token
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	// IDENT  = "IDENT"  // add, foobar, x, y, ...

	INT    = "INT"    // 1343456
	STRING = "STRING" // 1343456

	// Operators
	// ASSIGN   = "="
	// PLUS     = "+"
	// MINUS    = "-"
	// BANG     = "!"
	// ASTERISK = "*"
	// SLASH    = "/"
	// LT       = "<"
	// GT       = ">"
	// EQ       = "=="
	// NOT_EQ   = "!="
	// Delimiters
	// COMMA     = ","
	// SEMICOLON = ";"
	// COLON     = ":"
	LPAREN = "("
	RPAREN = ")"
	// LBRACE    = "{"
	// RBRACE    = "}"
	// LBRACKET  = "["
	// RBRACKET  = "]"

	// Keywords
	// FUNCTION = "FUNCTION"
	// LET      = "LET"
	TRUE  = "TRUE"
	FALSE = "FALSE"
	NIL   = "NIL"
	// IF       = "IF"
	// ELSE     = "ELSE"
	// RETURN   = "RETURN"

	SYMBOL = "SYMBOL"
)

var keywords = map[string]TokenType{
	// "fn":     FUNCTION,
	// "let":    LET,
	// "if":     IF,
	// "return": RETURN,
	"true":  TRUE,
	"false": FALSE,
	"nil":   NIL,
	// "else":   ELSE,
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readStringSymbol()
			t, ok := keywords[tok.Literal]
			if !ok {
				tok.Type = SYMBOL
			} else {
				tok.Type = t
			}
			return tok
		}
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = INT
			return tok
		} else {
			tok = newToken(SYMBOL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// indexing the input this way assumes all characters in the input
		// are ASCII, or the same number of bytes. Therefore this lexer would
		// not fully support Unicode
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readStringSymbol() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // slice out the identifier string from the input
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
