package introspection

import (
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

type __Type struct {
	schema *ast.Schema
	def    *ast.Definition
	typ    *ast.Type
}

func wrapTypeFromDef(s *ast.Schema, d *ast.Definition) *__Type {
	if d == nil {
		return nil
	}
	return &__Type{schema: s, def: d}
}

func wrapTypeFromType(s *ast.Schema, t *ast.Type) *__Type {
	if t == nil {
		return nil
	}

	if !t.NonNull && t.NamedType != "" {
		return &__Type{schema: s, def: s.Types[t.NamedType]}
	}
	return &__Type{schema: s, typ: t}
}

func defaultValue(value *ast.Value) *string {
	if value == nil {
		return nil
	}
	val := value.String()
	return &val
}

func (my *__Type) Kind() string {
	if my.typ != nil {
		if my.typ.NonNull {
			return "NON_NULL"
		}

		if my.typ.Elem != nil {
			return "LIST"
		}
	} else {
		return string(my.def.Kind)
	}

	panic("UNKNOWN")
}

func (my *__Type) Name() *string {
	if my.def == nil {
		return nil
	}
	return &my.def.Name
}

func (my *__Type) Description() *string {
	if my.def == nil || my.def.Description == "" {
		return nil
	}
	return &my.def.Description
}

func (my *__Type) Fields(includeDeprecated bool) []__Field {
	if my.def == nil || (my.def.Kind != ast.Object && my.def.Kind != ast.Interface) {
		return []__Field{}
	}
	var fields []__Field
	for _, f := range my.def.Fields {
		if strings.HasPrefix(f.Name, "__") {
			continue
		}

		if !includeDeprecated && f.Directives.ForName("deprecated") != nil {
			continue
		}

		var args []__InputValue
		for _, arg := range f.Arguments {
			args = append(args, __InputValue{
				Type:         wrapTypeFromType(my.schema, arg.Type),
				Name:         arg.Name,
				description:  arg.Description,
				DefaultValue: defaultValue(arg.DefaultValue),
			})
		}

		fields = append(fields, __Field{
			Name:        f.Name,
			description: f.Description,
			Args:        args,
			Type:        wrapTypeFromType(my.schema, f.Type),
			deprecation: f.Directives.ForName("deprecated"),
		})
	}
	return fields
}

func (my *__Type) InputFields() []__InputValue {
	if my.def == nil || my.def.Kind != ast.InputObject {
		return []__InputValue{}
	}

	var res []__InputValue
	for _, f := range my.def.Fields {
		res = append(res, __InputValue{
			Name:         f.Name,
			description:  f.Description,
			Type:         wrapTypeFromType(my.schema, f.Type),
			DefaultValue: defaultValue(f.DefaultValue),
		})
	}
	return res
}

func (my *__Type) Interfaces() []__Type {
	if my.def == nil || my.def.Kind != ast.Object {
		return []__Type{}
	}

	var res []__Type
	for _, i := range my.def.Interfaces {
		res = append(res, *wrapTypeFromDef(my.schema, my.schema.Types[i]))
	}

	return res
}

func (my *__Type) PossibleTypes() []__Type {
	if my.def == nil || (my.def.Kind != ast.Interface && my.def.Kind != ast.Union) {
		return []__Type{}
	}

	var res []__Type
	for _, pt := range my.schema.GetPossibleTypes(my.def) {
		res = append(res, *wrapTypeFromDef(my.schema, pt))
	}
	return res
}

func (my *__Type) EnumValues(includeDeprecated bool) []__EnumValue {
	if my.def == nil || my.def.Kind != ast.Enum {
		return []__EnumValue{}
	}

	var res []__EnumValue
	for _, val := range my.def.EnumValues {
		if !includeDeprecated && val.Directives.ForName("deprecated") != nil {
			continue
		}

		res = append(res, __EnumValue{
			Name:        val.Name,
			description: val.Description,
			deprecation: val.Directives.ForName("deprecated"),
		})
	}
	return res
}

func (my *__Type) OfType() *__Type {
	if my.typ == nil {
		return nil
	}
	if my.typ.NonNull {
		// fake non null nodes
		cpy := *my.typ
		cpy.NonNull = false

		return wrapTypeFromType(my.schema, &cpy)
	}
	return wrapTypeFromType(my.schema, my.typ.Elem)
}

func (my *__Type) SpecifiedByURL() *string {
	directive := my.def.Directives.ForName("specifiedBy")
	if my.def.Kind != ast.Scalar || directive == nil {
		return nil
	}
	// def: directive @specifiedBy(url: String!) on SCALAR
	// the argument "url" is required.
	url := directive.Arguments.ForName("url")
	return &url.Value.Raw
}
