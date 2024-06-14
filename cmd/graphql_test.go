package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"testing"
)

func TestGraphql(t *testing.T) {
	User := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "test", nil
				},
			},
		},
	})
	// Schema
	fields := graphql.Fields{
		"user": &graphql.Field{
			Type: User,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				p.Info.Operation.GetKind()
				return map[string]any{
					"name": "foo",
				}, nil
			},
		},
	}
	root := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	config := graphql.SchemaConfig{Query: graphql.NewObject(root)}
	schema, err := graphql.NewSchema(config)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			user {
				name
			}
		}
	`
	param := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(param)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	data, _ := json.Marshal(r)
	fmt.Printf("%s \n", data) // {"data":{"user":{"name":"foo"}}}
}
