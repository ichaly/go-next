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

	res["name"] = my.d.Name
	if len(my.d.Description) > 0 {
		res["description"] = my.d.Description
	}
	if !strings.HasPrefix(my.d.Type.Name(), "__") {
		res["type"] = &__Type{t: my.d.Type}
	}

	//必须存在不能为nil
	args := make([]__InputValue, 0, len(my.d.Arguments))
	for _, a := range my.d.Arguments {
		args = append(args, __InputValue{s: my.s, d: a})
	}
	res["args"] = args

	directive := my.d.Directives.ForName("deprecated")
	res["isDeprecated"] = directive != nil
	if directive != nil {
		reason := directive.Arguments.ForName("reason")
		if reason != nil && reason.Value != nil {
			res["deprecationReason"] = reason.Value.Raw
		}
	}

	return json.Marshal(res)
}
