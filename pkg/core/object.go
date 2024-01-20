package core

import (
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"reflect"
)

type GqlDescription interface {
	Description() string
}

func (my *Engine) asObject(typ reflect.Type) (graphql.Type, error) {
	typ, err := unwrap(typ)
	if err != nil {
		return nil, err
	}

	if obj, ok := my.types[typ.Name()]; ok {
		return obj.(*graphql.Object), nil
	}

	name, desc := typ.Name(), ""
	if v, ok := isImplements[GqlDescription](typ); ok {
		desc = v.Description()
	}

	o := graphql.NewObject(graphql.ObjectConfig{
		Name: name, Description: desc, Fields: graphql.Fields{},
	})
	if err = my.parseFields(typ, o); err != nil {
		return nil, err
	}

	my.types[name] = o
	return o, nil
}

func (my *Engine) parseFields(typ reflect.Type, obj *graphql.Object) error {
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		// 忽略私有字段
		if !f.IsExported() {
			continue
		}
		// 递归匿名字段
		if f.Anonymous {
			if sub, err := unwrap(f.Type); err != nil {
				return err
			} else if err = my.parseFields(sub, obj); err != nil {
				return err
			}
			continue
		}

		fieldType, err := parseType(f.Type, "obj field",
			my.asBuiltinScalar, my.asId, my.asObject,
		)
		if err != nil {
			return err
		}

		fieldName := strcase.ToLowerCamel(f.Name)
		obj.AddFieldConfig(fieldName, &graphql.Field{
			Type: wrapType(f.Type, fieldType),
		})
	}
	return nil
}
