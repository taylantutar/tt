package lexer

import (
	"strings"
	"unicode"
	"github.com/taylantutar/tt/pkg/parser"
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

func (l *Lexer) NextToken() parser.Token {
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok := parser.Token{Type: parser.ASSIGN, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '+':
		tok := parser.Token{Type: parser.PLUS, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '-':
		tok := parser.Token{Type: parser.MINUS, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '*':
		tok := parser.Token{Type: parser.ASTERISK, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '/':
		tok := parser.Token{Type: parser.SLASH, Literal: string(l.ch)}
		l.readChar()
		return tok
	case 0:
		return parser.Token{Type: parser.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			return parser.Token{Type: lookupIdent(literal), Literal: literal}
		} else if isDigit(l.ch) {
			return parser.Token{Type: parser.INT, Literal: l.readNumber()}
		} else {
			tok := parser.Token{Type: parser.ILLEGAL, Literal: string(l.ch)}
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

func lookupIdent(ident string) parser.TokenType {
	switch strings.ToLower(ident) {
	case "set":
		return parser.SET
	case "print":
		return parser.PRINT
	default:
		return parser.IDENT
	}
}
