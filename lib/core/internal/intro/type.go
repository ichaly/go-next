package intro

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type __FullType struct {
	s *ast.Schema
	d *ast.Definition
}

func (my __FullType) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	res["name"] = my.d.Name
	res["kind"] = my.d.Kind
	if len(my.d.Description) > 0 {
		res["description"] = my.d.Description
	}
	if len(my.d.Fields) > 0 {
		fields := make([]__Field, 0, len(my.d.Fields))
		for _, f := range my.d.Fields {
			if !strings.HasPrefix(f.Name, "__") {
				fields = append(fields, __Field{s: my.s, d: f})
			}
		}
		res["fields"] = fields
	}
	if my.d.Kind == ast.Object || my.d.Kind == ast.Interface {
		//如果是Object,必须存在不能为nil
		res["interfaces"] = []interface{}{}
		for _, v := range my.d.Interfaces {
			res["interfaces"] = &__FullType{s: my.s, d: my.s.Types[v]}
		}
	}

	return json.Marshal(res)
}

type __RootType struct {
	d *ast.Definition
}

func (my __RootType) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})
	res["name"] = my.d.Name
	return json.Marshal(res)
}

type __Type struct {
	t *ast.Type
}

func (my __Type) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})

	res["name"] = my.t.Name()
	if my.t.NonNull {
		res["kind"] = "NON_NULL"
	}

	//if my.t.Elem != nil {
	//	res["kind"] = "List"
	//}
	//res["ofType"] = &__Type{t: my.t.Elem}

	return json.Marshal(res)
}
