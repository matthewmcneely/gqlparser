package ast

type DefinitionKind string

const (
	Scalar      DefinitionKind = "SCALAR"
	Object      DefinitionKind = "OBJECT"
	Interface   DefinitionKind = "INTERFACE"
	Union       DefinitionKind = "UNION"
	Enum        DefinitionKind = "ENUM"
	InputObject DefinitionKind = "INPUT_OBJECT"
)

// Definition is the core type definition object, it includes all of the definable types
// but does *not* cover schema or directives.
//
// @vektah: Javascript implementation has different types for all of these, but they are
// more similar than different and don't define any behaviour. I think this style of
// "some hot" struct works better, at least for go.
//
// Type extensions are also represented by this same struct.
type Definition struct {
	Kind        DefinitionKind `json:"kind,omitempty"`
	Description string         `json:"description,omitempty"`
	Name        string         `json:"name,omitempty"`
	Directives  DirectiveList  `json:"directives,omitempty"`
	Interfaces  []string       `json:"interfaces,omitempty"` // object and input object
	Fields      FieldList      `json:"fields,omitempty"`     // object and input object
	Types       []string       `json:"types,omitempty"`      // union
	EnumValues  EnumValueList  `json:"enumValues,omitempty"` // enum

	Position *Position `dump:"-" json:"-"`
	BuiltIn  bool      `dump:"-" json:"-"`
}

func (d *Definition) IsLeafType() bool {
	return d.Kind == Enum || d.Kind == Scalar
}

func (d *Definition) IsAbstractType() bool {
	return d.Kind == Interface || d.Kind == Union
}

func (d *Definition) IsCompositeType() bool {
	return d.Kind == Object || d.Kind == Interface || d.Kind == Union
}

func (d *Definition) IsInputType() bool {
	return d.Kind == Scalar || d.Kind == Enum || d.Kind == InputObject
}

func (d *Definition) OneOf(types ...string) bool {
	for _, t := range types {
		if d.Name == t {
			return true
		}
	}
	return false
}

type FieldDefinition struct {
	Description  string                 `json:"description,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Arguments    ArgumentDefinitionList `json:"arguments,omitempty"`    // only for objects
	DefaultValue *Value                 `json:"defaultValue,omitempty"` // only for input objects
	Type         *Type                  `json:"type,omitempty"`
	Directives   DirectiveList          `json:"directives,omitempty"`
	Position     *Position              `dump:"-" json:"-"`
}

type ArgumentDefinition struct {
	Description  string        `json:"description,omitempty"`
	Name         string        `json:"name,omitempty"`
	DefaultValue *Value        `json:"defaultValue,omitempty"`
	Type         *Type         `json:"type,omitempty"`
	Directives   DirectiveList `json:"directives,omitempty"`
	Position     *Position     `dump:"-" json:"-"`
}

type EnumValueDefinition struct {
	Description string        `json:"description,omitempty"`
	Name        string        `json:"name,omitempty"`
	Directives  DirectiveList `json:"directives,omitempty"`
	Position    *Position     `dump:"-" json:"-"`
}

type DirectiveDefinition struct {
	Description  string                 `json:"description,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Arguments    ArgumentDefinitionList `json:"arguments,omitempty"`
	Locations    []DirectiveLocation    `json:"locations,omitempty"`
	IsRepeatable bool                   `json:"isRepeatable,omitempty"`
	Position     *Position              `dump:"-" json:"-"`
}
