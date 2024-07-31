package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type __Field struct {
	s *ast.Schema
	d *ast.FieldDefinition
}

func (my __Field) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if my.d != nil {
		res["name"] = my.d.Name
		if len(my.d.Description) > 0 {
			res["description"] = my.d.Description
		}
		if !strings.HasPrefix(my.d.Type.Name(), "__") {
			res["type"] = &__Type{s: my.s, d: my.s.Types[my.d.Type.Name()]}
		}
	}

	return json.Marshal(res)
}
