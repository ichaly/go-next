// introspection implements the spec defined in https://github.com/facebook/graphql/blob/master/spec/Section%204%20--%20Introspection.md#schema-introspection
package introspection

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type (
	__Field struct {
		Name        string
		Description string
		Type        *__Type
		Args        []__InputValue
		Deprecation *ast.Directive
	}

	__Directive struct {
		Name         string
		Description  string
		Locations    []string
		Args         []__InputValue
		IsRepeatable bool
	}

	__InputValue struct {
		Name         string
		Description  string
		DefaultValue *string
		Type         *__Type
		Deprecation  *ast.Directive
	}

	__EnumValue struct {
		Name        string
		Description string
		Deprecation *ast.Directive
	}
)

func (my __Field) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if my.Name != "" {
		res["name"] = my.Name
	}
	if my.Description != "" {
		res["description"] = my.Description
	}
	if my.Type != nil {
		res["type"] = my.Type
	}
	if len(my.Args) > 0 {
		res["args"] = my.Args
	}
	res["isDeprecated"] = my.Deprecation != nil
	if my.Deprecation != nil {
		reason := my.Deprecation.Arguments.ForName("reason")
		if reason == nil {
			res["deprecationReason"] = reason.Value.Raw
		}
	}

	return json.Marshal(res)
}

func (my __Directive) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if my.Name != "" {
		res["name"] = my.Name
	}
	if my.Description != "" {
		res["description"] = my.Description
	}
	res["isRepeatable"] = my.IsRepeatable
	if len(my.Locations) > 0 {
		res["locations"] = my.Locations
	}
	if len(my.Args) > 0 {
		res["args"] = my.Args
	}

	return json.Marshal(res)
}

func (my __InputValue) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if my.Name != "" {
		res["name"] = my.Name
	}
	if my.Description != "" {
		res["description"] = my.Description
	}
	if my.Type != nil {
		res["type"] = my.Type
	}
	res["isDeprecated"] = my.Deprecation != nil
	if my.Deprecation != nil {
		reason := my.Deprecation.Arguments.ForName("reason")
		if reason == nil {
			res["deprecationReason"] = reason.Value.Raw
		}
	}
	if my.DefaultValue != nil {
		res["defaultValue"] = my.DefaultValue
	}

	return json.Marshal(res)
}

func (my __EnumValue) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if my.Name != "" {
		res["name"] = my.Name
	}
	if my.Description != "" {
		res["description"] = my.Description
	}
	res["isDeprecated"] = my.Deprecation != nil
	if my.Deprecation != nil {
		reason := my.Deprecation.Arguments.ForName("reason")
		if reason == nil {
			res["deprecationReason"] = reason.Value.Raw
		}
	}

	return json.Marshal(res)
}
