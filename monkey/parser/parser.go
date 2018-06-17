// parser/parser.go

package parser

import (
	"github.com/515hikaru/monkey-go/monkey/ast"
	"github.com/515hikaru/monkey-go/monkey/lexer"
	"github.com/515hikaru/monkey-go/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つのトークンを読み込む。 curToken と peekToken の両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
