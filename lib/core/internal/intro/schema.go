package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type __Schema struct {
	s *ast.Schema
}

func New(s *ast.Schema) __Schema {
	return __Schema{s: s}
}

func (my __Schema) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if len(my.s.Types) > 0 {
		types := make(map[string]__Type, len(my.s.Types))
		for k, v := range my.s.Types {
			types[k] = __Type{s: my.s, d: v}
		}
		res["types"] = types
	}
	if my.s.Query != nil {
		res["queryType"] = __Type{s: my.s, d: my.s.Query}
	}

	return json.Marshal(res)
}
