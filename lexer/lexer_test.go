package lexer

import (
	"testing"

	"github.com/matthewmcneely/gqlparser/gqlerror"

	"github.com/matthewmcneely/gqlparser/ast"
	"github.com/matthewmcneely/gqlparser/parser/testrunner"
)

func TestLexer(t *testing.T) {
	testrunner.Test(t, "lexer_test.yml", func(t *testing.T, input string) testrunner.Spec {
		l := New(&ast.Source{Input: input, Name: "spec"})

		ret := testrunner.Spec{}
		for {
			tok, err := l.ReadToken()

			if err != nil {
				ret.Error = err.(*gqlerror.Error)
				break
			}

			if tok.Kind == EOF {
				break
			}

			ret.Tokens = append(ret.Tokens, testrunner.Token{
				Kind:   tok.Kind.Name(),
				Value:  tok.Value,
				Line:   tok.Pos.Line,
				Column: tok.Pos.Column,
				Start:  tok.Pos.Start,
				End:    tok.Pos.End,
				Src:    tok.Pos.Src.Name,
			})
		}

		return ret
	})
}
