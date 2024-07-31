package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type __InputValue struct {
	s *ast.Schema
	d *ast.ArgumentDefinition
}

func (my __InputValue) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if len(my.d.Name) > 0 {
		res["name"] = my.d.Name
	}
	if len(my.d.Description) > 0 {
		res["description"] = my.d.Description
	}
	if my.d.Type != nil {
		res["type"] = __Type{s: my.s, d: my.s.Types[my.d.Type.Name()]}
	}

	return json.Marshal(res)
}
