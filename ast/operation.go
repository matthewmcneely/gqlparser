package ast

type Operation string

const (
	Query        Operation = "query"
	Mutation     Operation = "mutation"
	Subscription Operation = "subscription"
)

type OperationDefinition struct {
	Operation           Operation              `json:"operation,omitempty"`
	Name                string                 `json:"name,omitempty"`
	VariableDefinitions VariableDefinitionList `json:"variableDefinitions,omitempty"`
	Directives          DirectiveList          `json:"directives,omitempty"`
	SelectionSet        SelectionSet           `json:"selectionSet,omitempty"`
	Position            *Position              `dump:"-" json:"-"`
}

type VariableDefinition struct {
	Variable     string        `json:"variable,omitempty"`
	Type         *Type         `json:"type,omitempty"`
	DefaultValue *Value        `json:"defaultValue,omitempty"`
	Directives   DirectiveList `json:"directives,omitempty"`
	Position     *Position     `dump:"-" json:"-"`

	// Requires validation
	Definition *Definition `json:"definition,omitempty"`
	Used       bool        `dump:"-" json:"used,omitempty"`
}
