package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	if tok = l.delimiter(); tok.Type != token.ILLEGAL {
		l.readChar()
		return tok
	}

	if tok = l.mathOperator(); tok.Type != token.ILLEGAL {
		l.readChar()
		return tok
	}

	if tok = l.relationalOperator(); tok.Type != token.ILLEGAL {
		l.readChar()
		return tok
	}

	if tok = l.logicalOperator(); tok.Type != token.ILLEGAL {
		l.readChar()
		return tok
	}

	switch l.ch {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) delimiter() token.Token {
	delimiters := map[byte]token.TokenType{
		';': token.SEMICOLON,
		',': token.COMMA,
		'(': token.LPAREN,
		')': token.RPAREN,
		'{': token.LBRACE,
		'}': token.RBRACE,
		'[': token.LBRACKET,
		']': token.RBRACKET,
	}

	if tokType, ok := delimiters[l.ch]; ok {
		return newToken(tokType, l.ch)
	}
	return newToken(token.ILLEGAL, l.ch)
}

func (l *Lexer) mathOperator() token.Token {
	mathOperators := map[byte]token.TokenType{
		'+': token.PLUS,
		'-': token.MINUS,
		'*': token.ASTERISK,
		'/': token.SLASH,
	}

	if tokType, ok := mathOperators[l.ch]; ok {
		return newToken(tokType, l.ch)
	}

	return newToken(token.ILLEGAL, l.ch)
}

func (l *Lexer) relationalOperator() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '<':
		tok = newToken(token.LT, l.ch)

	case '>':
		tok = newToken(token.GT, l.ch)

	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}
	return tok
}

func (l *Lexer) logicalOperator() token.Token {
	if l.ch == '!' && l.peekChar() != '=' {
		return newToken(token.BANG, l.ch)
	}
	return newToken(token.ILLEGAL, l.ch)
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
