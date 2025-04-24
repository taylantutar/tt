package lexer

import (
	"github.com/taylantutar/tt/pkg/token"
	"strings"
	"unicode"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // null karakter
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		tok := token.Token{Type: token.ASSIGN, Literal: string(l.currentChar)}
		l.readChar()
		return tok
	case '+':
		tok := token.Token{Type: token.PLUS, Literal: string(l.currentChar)}
		l.readChar()
		return tok
	case '-':
		tok := token.Token{Type: token.MINUS, Literal: string(l.currentChar)}
		l.readChar()
		return tok
	case '*':
		tok := token.Token{Type: token.ASTERISK, Literal: string(l.currentChar)}
		l.readChar()
		return tok
	case '/':
		tok := token.Token{Type: token.SLASH, Literal: string(l.currentChar)}
		l.readChar()
		return tok
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.currentChar) {
			literal := l.readIdentifier()
			return token.Token{Type: lookupIdent(literal), Literal: literal}
		} else if isDigit(l.currentChar) {
			return token.Token{Type: token.INT, Literal: l.readNumber()}
		} else {
			tok := token.Token{Type: token.ILLEGAL, Literal: string(l.currentChar)}
			l.readChar()
			return tok
		}
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\t' || l.currentChar == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func lookupIdent(ident string) token.TokenType {
	switch strings.ToLower(ident) {
	case "set":
		return token.SET
	case "print":
		return token.PRINT
	default:
		return token.IDENT
	}
}
