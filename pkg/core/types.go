package core

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"reflect"
)

type SchemaMeta[P any, R any] struct {
	parentType P
	resultType R
}

type Schema interface {
	Resolve(p graphql.ResolveParams) (interface{}, error)
}

type (
	Query        struct{}
	Mutation     struct{}
	Subscription struct{}
)

type typeParser func(typ reflect.Type) (graphql.Type, error)

var (
	Void = graphql.NewScalar(graphql.ScalarConfig{
		Name:         "Void",
		Description:  "void",
		Serialize:    func(value interface{}) interface{} { return "0" },
		ParseValue:   func(value interface{}) interface{} { return 0 },
		ParseLiteral: func(valueAST ast.Value) interface{} { return 0 },
	})

	Cursor = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "Cursor",
		Description: "A `Cursor` is an encoded string use for pagination",
		Serialize: func(val interface{}) interface{} {
			//js := []byte(fmt.Sprintf(`{ me: "null", posts_cursor: "%v12345" }`, encPrefix))
			//nonce := sha256.Sum256(js)
			//out, err := encryptValues(js, []byte(encPrefix), []byte(decPrefix), nonce[:], key)
			//if err != nil {
			//	return nil
			//}
			return val
		},
		ParseValue: func(val interface{}) interface{} {
			return val
		},
		ParseLiteral: func(val ast.Value) interface{} {
			return nil
		},
	})

	Json = graphql.NewScalar(
		graphql.ScalarConfig{
			Name:        "Json",
			Description: "The `Json` scalar type represents Json values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf)",
			Serialize: func(value interface{}) interface{} {
				return value
			},
			ParseValue: func(value interface{}) interface{} {
				return value
			},
			ParseLiteral: parseLiteral,
		},
	)
)

func parseLiteral(astValue ast.Value) interface{} {
	kind := astValue.GetKind()
	switch kind {
	case kinds.StringValue:
		return astValue.GetValue()
	case kinds.BooleanValue:
		return astValue.GetValue()
	case kinds.IntValue:
		return astValue.GetValue()
	case kinds.FloatValue:
		return astValue.GetValue()
	case kinds.ObjectValue:
		obj := make(map[string]interface{})
		for _, v := range astValue.GetValue().([]*ast.ObjectField) {
			obj[v.Name.Value] = parseLiteral(v.Value)
		}
		return obj
	case kinds.ListValue:
		list := make([]interface{}, 0)
		for _, v := range astValue.GetValue().([]ast.Value) {
			list = append(list, parseLiteral(v))
		}
		return list
	default:
		return nil
	}
}
