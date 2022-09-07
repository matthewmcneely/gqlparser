package ast

type FragmentSpread struct {
	Name       string        `json:"name,omitempty"`
	Directives DirectiveList `json:"directives,omitempty"`

	// Require validation
	ObjectDefinition *Definition         `json:"objectDefinition,omitempty"`
	Definition       *FragmentDefinition `json:"definition,omitempty"`

	Position *Position `dump:"-" json:"-"`
}

type InlineFragment struct {
	TypeCondition string        `json:"typeCondition,omitempty"`
	Directives    DirectiveList `json:"directives,omitempty"`
	SelectionSet  SelectionSet  `json:"selectionSet,omitempty"`

	// Require validation
	ObjectDefinition *Definition `json:"objectDefinition,omitempty"`

	Position *Position `dump:"-" json:"-"`
}

type FragmentDefinition struct {
	Name string `json:"name,omitempty"`
	// Note: fragment variable definitions are experimental and may be changed
	// or removed in the future.
	VariableDefinition VariableDefinitionList `json:"variableDefinition,omitempty"`
	TypeCondition      string                 `json:"typeCondition,omitempty"`
	Directives         DirectiveList          `json:"directives,omitempty"`
	SelectionSet       SelectionSet           `json:"selectionSet,omitempty"`

	// Require validation
	Definition *Definition `json:"definition,omitempty"`

	Position *Position `dump:"-" json:"-"`
}
