package parser

import (
	"github.com/matthewmcneely/gqlparser/gqlerror"
	"testing"

	"github.com/matthewmcneely/gqlparser/ast"
	"github.com/matthewmcneely/gqlparser/parser/testrunner"
)

func TestQueryDocument(t *testing.T) {
	testrunner.Test(t, "query_test.yml", func(t *testing.T, input string) testrunner.Spec {
		doc, err := ParseQuery(&ast.Source{Input: input, Name: "spec"})
		if err != nil {
			gqlErr := err.(*gqlerror.Error)
			return testrunner.Spec{
				Error: gqlErr,
				AST:   ast.Dump(doc),
			}
		}
		return testrunner.Spec{
			AST: ast.Dump(doc),
		}
	})
}
