package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type __Field struct {
	s *ast.Schema
	d *ast.FieldDefinition
}

func (my __Field) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	res["name"] = my.d.Name
	res["description"] = my.d.Description
	res["type"] = &__Type{s: my.s, d: my.s.Types[my.d.Type.NamedType]}

	return json.Marshal(res)
}
