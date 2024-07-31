package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type __Type struct {
	s *ast.Schema
	d *ast.Definition
}

func (my __Type) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	res["name"] = my.d.Name
	if len(my.d.Kind) > 0 {
		res["kind"] = my.d.Kind
	}
	if len(my.d.Description) > 0 {
		res["description"] = my.d.Description
	}
	if len(my.d.Fields) > 0 {
		fields := make([]__Field, 0, len(my.d.Fields))
		for _, f := range my.d.Fields {
			fields = append(fields, __Field{s: my.s, d: f})
		}
		res["fields"] = fields
	}

	return json.Marshal(res)
}
