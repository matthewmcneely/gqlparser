package ast

type SelectionSet []Selection

type Selection interface {
	isSelection()
	GetPosition() *Position
}

func (*Field) isSelection()          {}
func (*FragmentSpread) isSelection() {}
func (*InlineFragment) isSelection() {}

func (s *Field) GetPosition() *Position          { return s.Position }
func (s *FragmentSpread) GetPosition() *Position { return s.Position }
func (s *InlineFragment) GetPosition() *Position { return s.Position }

type Field struct {
	Alias        string        `json:"alias,omitempty"`
	Name         string        `json:"name,omitempty"`
	Arguments    ArgumentList  `json:"arguments,omitempty"`
	Directives   DirectiveList `json:"directives,omitempty"`
	SelectionSet SelectionSet  `json:"selectionSet,omitempty"`
	Position     *Position     `dump:"-" json:"-"`

	// Require validation
	Definition       *FieldDefinition `json:"definition,omitempty"`
	ObjectDefinition *Definition      `json:"objectDefinition,omitempty"`
}

type Argument struct {
	Name     string    `json:"name,omitempty"`
	Value    *Value    `json:"value,omitempty"`
	Position *Position `dump:"-" json:"-"`
}

func (f *Field) ArgumentMap(vars map[string]interface{}) map[string]interface{} {
	return arg2map(f.Definition.Arguments, f.Arguments, vars)
}
