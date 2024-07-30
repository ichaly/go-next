package introspection

import (
	"encoding/json"
	"sort"

	"github.com/vektah/gqlparser/v2/ast"
)

type __Schema struct {
	schema *ast.Schema
}

func NewIntrospection(s *ast.Schema) __Schema {
	return __Schema{schema: s}
}

func (my __Schema) Types() []__Type {
	typeIndex := map[string]__Type{}
	typeNames := make([]string, 0, len(my.schema.Types))
	for _, t := range my.schema.Types {
		typeNames = append(typeNames, t.Name)
		typeIndex[t.Name] = *wrapTypeFromDef(my.schema, t)
	}
	sort.Strings(typeNames)

	types := make([]__Type, len(typeNames))
	for i, t := range typeNames {
		types[i] = typeIndex[t]
	}
	return types
}

func (my __Schema) QueryType() *__Type {
	return wrapTypeFromDef(my.schema, my.schema.Query)
}

func (my __Schema) MutationType() *__Type {
	return wrapTypeFromDef(my.schema, my.schema.Mutation)
}

func (my __Schema) SubscriptionType() *__Type {
	return wrapTypeFromDef(my.schema, my.schema.Subscription)
}

func (my __Schema) Directives() []__Directive {
	dIndex := map[string]__Directive{}
	dNames := make([]string, 0, len(my.schema.Directives))

	for _, d := range my.schema.Directives {
		dNames = append(dNames, d.Name)
		dIndex[d.Name] = my.directiveFromDef(d)
	}
	sort.Strings(dNames)

	res := make([]__Directive, len(dNames))
	for i, d := range dNames {
		res[i] = dIndex[d]
	}

	return res
}

func (my __Schema) Description() *string {
	if my.schema.Description == "" {
		return nil
	}
	return &my.schema.Description
}

func (my __Schema) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	if my.Description() != nil {
		res["description"] = my.Description()
	}
	if len(my.Directives()) > 0 {
		res["directives"] = my.Directives()
	}
	if len(my.Types()) > 0 {
		res["types"] = my.Types()
	}
	if my.QueryType() != nil {
		res["queryType"] = my.QueryType()
	}
	if my.MutationType() != nil {
		res["mutationType"] = my.MutationType()
	}
	if my.SubscriptionType() != nil {
		res["subscriptionType"] = my.SubscriptionType()
	}

	return json.Marshal(res)
}

func (my __Schema) directiveFromDef(d *ast.DirectiveDefinition) __Directive {
	locs := make([]string, len(d.Locations))
	for i, l := range d.Locations {
		locs[i] = string(l)
	}

	args := make([]__InputValue, len(d.Arguments))
	for i, a := range d.Arguments {
		args[i] = __InputValue{
			Name:         a.Name,
			Description:  a.Description,
			DefaultValue: defaultValue(a.DefaultValue),
			Type:         wrapTypeFromType(my.schema, a.Type),
		}
	}

	return __Directive{
		Name:         d.Name,
		Args:         args,
		Locations:    locs,
		Description:  d.Description,
		IsRepeatable: d.IsRepeatable,
	}
}
