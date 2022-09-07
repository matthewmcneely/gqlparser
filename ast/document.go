package ast

type QueryDocument struct {
	Operations OperationList          `json:"operations,omitempty"`
	Fragments  FragmentDefinitionList `json:"fragments,omitempty"`
	Position   *Position              `dump:"-" json:"-"`
}

type SchemaDocument struct {
	Schema          SchemaDefinitionList    `json:"schema,omitempty"`
	SchemaExtension SchemaDefinitionList    `json:"schemaExtension,omitempty"`
	Directives      DirectiveDefinitionList `json:"directives,omitempty"`
	Definitions     DefinitionList          `json:"definitions,omitempty"`
	Extensions      DefinitionList          `json:"extensions,omitempty"`
	Position        *Position               `dump:"-" json:"-"`
}

func (d *SchemaDocument) Merge(other *SchemaDocument) {
	d.Schema = append(d.Schema, other.Schema...)
	d.SchemaExtension = append(d.SchemaExtension, other.SchemaExtension...)
	d.Directives = append(d.Directives, other.Directives...)
	d.Definitions = append(d.Definitions, other.Definitions...)
	d.Extensions = append(d.Extensions, other.Extensions...)
}

type Schema struct {
	Query        *Definition
	Mutation     *Definition
	Subscription *Definition

	Types      map[string]*Definition
	Directives map[string]*DirectiveDefinition

	PossibleTypes map[string][]*Definition
	Implements    map[string][]*Definition

	Description string
}

// AddTypes is the helper to add types definition to the schema
func (s *Schema) AddTypes(defs ...*Definition) {
	if s.Types == nil {
		s.Types = make(map[string]*Definition)
	}
	for _, def := range defs {
		s.Types[def.Name] = def
	}
}

func (s *Schema) AddPossibleType(name string, def *Definition) {
	s.PossibleTypes[name] = append(s.PossibleTypes[name], def)
}

// GetPossibleTypes will enumerate all the definitions for a given interface or union
func (s *Schema) GetPossibleTypes(def *Definition) []*Definition {
	return s.PossibleTypes[def.Name]
}

func (s *Schema) AddImplements(name string, iface *Definition) {
	s.Implements[name] = append(s.Implements[name], iface)
}

// GetImplements returns all the interface and union definitions that the given definition satisfies
func (s *Schema) GetImplements(def *Definition) []*Definition {
	return s.Implements[def.Name]
}

type SchemaDefinition struct {
	Description    string                      `json:"description,omitempty"`
	Directives     DirectiveList               `json:"directives,omitempty"`
	OperationTypes OperationTypeDefinitionList `json:"operationTypes,omitempty"`
	Position       *Position                   `dump:"-" json:"-"`
}

type OperationTypeDefinition struct {
	Operation Operation `json:"operation,omitempty"`
	Type      string    `json:"type,omitempty"`
	Position  *Position `dump:"-" json:"-"`
}
