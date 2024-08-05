package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type __EnumValue struct {
	d *ast.EnumValueDefinition
}

func (my __EnumValue) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if len(my.d.Name) > 0 {
		res["name"] = my.d.Name
	}
	if len(my.d.Description) > 0 {
		res["description"] = my.d.Description
	}
	deprecation := my.d.Directives.ForName("deprecated")
	res["isDeprecated"] = deprecation != nil
	if deprecation != nil {
		if reason := deprecation.Arguments.ForName("reason"); reason != nil {
			res["deprecationReason"] = reason.Value.Raw
		}
	}

	return json.Marshal(res)
}
