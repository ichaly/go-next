package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type __Directive struct {
	s *ast.Schema
	d *ast.DirectiveDefinition
}

func (my __Directive) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if len(my.d.Name) > 0 {
		res["name"] = my.d.Name
	}
	if len(my.d.Description) > 0 {
		res["description"] = my.d.Description
	}
	if len(my.d.Locations) > 0 {
		res["locations"] = my.d.Locations
	}
	args := make([]__InputValue, 0, len(my.d.Arguments))
	for _, v := range my.d.Arguments {
		args = append(args, __InputValue{s: my.s, d: v})
	}
	res["args"] = args

	return json.Marshal(res)
}
