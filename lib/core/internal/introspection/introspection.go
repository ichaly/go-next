// introspection implements the spec defined in https://github.com/facebook/graphql/blob/master/spec/Section%204%20--%20Introspection.md#schema-introspection
package introspection

import "github.com/vektah/gqlparser/v2/ast"

type (
	__Directive struct {
		Name         string
		description  string
		Locations    []string
		Args         []__InputValue
		IsRepeatable bool
	}

	__EnumValue struct {
		Name        string
		description string
		deprecation *ast.Directive
	}

	__Field struct {
		Name        string
		description string
		Type        *__Type
		Args        []__InputValue
		deprecation *ast.Directive
	}

	__InputValue struct {
		Name         string
		description  string
		DefaultValue *string
		Type         *__Type
	}
)

func (my *__EnumValue) Description() *string {
	if my.description == "" {
		return nil
	}
	return &my.description
}

func (my *__EnumValue) IsDeprecated() bool {
	return my.deprecation != nil
}

func (my *__EnumValue) DeprecationReason() *string {
	if my.deprecation == nil {
		return nil
	}

	reason := my.deprecation.Arguments.ForName("reason")
	if reason == nil {
		return nil
	}

	return &reason.Value.Raw
}

func (my *__Field) Description() *string {
	if my.description == "" {
		return nil
	}
	return &my.description
}

func (my *__Field) IsDeprecated() bool {
	return my.deprecation != nil
}

func (my *__Field) DeprecationReason() *string {
	if my.deprecation == nil || !my.IsDeprecated() {
		return nil
	}

	reason := my.deprecation.Arguments.ForName("reason")

	if reason == nil {
		defaultReason := "No longer supported"
		return &defaultReason
	}

	return &reason.Value.Raw
}

func (my *__InputValue) Description() *string {
	if my.description == "" {
		return nil
	}
	return &my.description
}

func (my *__Directive) Description() *string {
	if my.description == "" {
		return nil
	}
	return &my.description
}
