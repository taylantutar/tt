package lexer

import (
	"strings"
	"unicode"
	"github.com/taylantutar/tt/pkg/token"
)

type Lexer struct {
	input        string
	position     int  
	readPosition int  
	ch           byte 
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // null karakter
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok := token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '+':
		tok := token.Token{Type: token.PLUS, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '-':
		tok := token.Token{Type: token.MINUS, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '*':
		tok := token.Token{Type: token.ASTERISK, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '/':
		tok := token.Token{Type: token.SLASH, Literal: string(l.ch)}
		l.readChar()
		return tok
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			return token.Token{Type: lookupIdent(literal), Literal: literal}
		} else if isDigit(l.ch) {
			return token.Token{Type: token.INT, Literal: l.readNumber()}
		} else {
			tok := token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
			l.readChar()
			return tok
		}
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
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
