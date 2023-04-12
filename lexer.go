package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func parse() {
	yyDebug = 3
	line := "11+2"
	lex := NewLexer(line)
	if yyParse(lex) != 0 {
		fmt.Println(lex.err)
		return
	}
	fmt.Println("result =", lex.val)
}

type Lexer struct {
	scanner *bufio.Scanner

	val int
	err error
}

func NewLexer(s string) *Lexer {
	scanner := bufio.NewScanner(strings.NewReader(s))
	lexer := &Lexer{scanner: scanner}
	scanner.Split(lexer.split)
	return lexer
}

func (l *Lexer) Lex(lval *yySymType) int {
	if !l.scanner.Scan() {
		l.err = l.scanner.Err()
		if l.err != nil {
			return yyErrCode
		}
		return 0
	}

	return 0
}

func (l *Lexer) Error(s string) {
	if l.err == nil {
		l.err = fmt.Errorf("syntax error: %s", s)
	}
}

func (l *Lexer) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	r := bytes.NewReader(data)
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			return 0, nil, err
		}
		_ = c
	}
}
