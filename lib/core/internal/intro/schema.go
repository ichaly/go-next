package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type __Schema struct {
	s *ast.Schema
}

func New(s *ast.Schema) interface{} {
	res := make(map[string]interface{})
	res["__schema"] = __Schema{s: s}
	return res
}

func (my __Schema) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if len(my.s.Types) > 0 {
		types := make([]__Type, 0, len(my.s.Types))
		for k, v := range my.s.Types {
			if !strings.HasPrefix(k, "__") {
				types = append(types, __Type{s: my.s, d: v})
			}
		}
		res["types"] = types
	}
	if my.s.Query != nil {
		d := &ast.Definition{Name: my.s.Query.Name}
		res["queryType"] = __Type{s: my.s, d: d}
	}

	return json.Marshal(res)
}
