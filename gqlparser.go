package gqlparser

import (
	"github.com/matthewmcneely/gqlparser/ast"
	"github.com/matthewmcneely/gqlparser/gqlerror"
	"github.com/matthewmcneely/gqlparser/parser"
	"github.com/matthewmcneely/gqlparser/validator"
	_ "github.com/matthewmcneely/gqlparser/validator/rules"
)

func LoadSchema(str ...*ast.Source) (*ast.Schema, error) {
	return validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
}

func MustLoadSchema(str ...*ast.Source) *ast.Schema {
	s, err := validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
	if err != nil {
		panic(err)
	}
	return s
}

func LoadQuery(schema *ast.Schema, str string) (*ast.QueryDocument, gqlerror.List) {
	query, err := parser.ParseQuery(&ast.Source{Input: str})
	if err != nil {
		gqlErr := err.(*gqlerror.Error)
		return nil, gqlerror.List{gqlErr}
	}
	errs := validator.Validate(schema, query)
	if errs != nil {
		return nil, errs
	}

	return query, nil
}

func MustLoadQuery(schema *ast.Schema, str string) *ast.QueryDocument {
	q, err := LoadQuery(schema, str)
	if err != nil {
		panic(err)
	}
	return q
}
