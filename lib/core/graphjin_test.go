package core

import (
	"database/sql"
	"github.com/dosco/graphjin/core/v3"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/core/internal/intro"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"net/http"
	"os"
	"testing"
)

type _GraphJinSuite struct {
	suite.Suite
	db *sql.DB
}

func TestGraphJin(t *testing.T) {
	suite.Run(t, new(_GraphJinSuite))
}

func (my *_GraphJinSuite) SetupSuite() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5678/demo?sslmode=disable")
	my.Require().NoError(err)
	my.db = db
}

func (my *_GraphJinSuite) TestGraphJin() {
	gj, err := core.NewGraphJin(&core.Config{
		EnableCamelcase: true,
		DisableAgg:      true,
		DisableFuncs:    true,
	}, my.db)
	my.Require().NoError(err)

	r := gin.Default()
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql", func(ctx *gin.Context) {
		var req struct {
			Query     string                 `form:"query"`
			Operation string                 `form:"operationName" json:"operationName"`
			Variables map[string]interface{} `form:"variables"`
		}
		_ = ctx.ShouldBindBodyWith(&req, binding.JSON)
		res, _ := gj.GraphQL(ctx, req.Query, nil, nil)
		ctx.JSON(http.StatusOK, res)
	})
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql0", func(ctx *gin.Context) {
		file, _ := os.ReadFile("./assets/schema.gql")
		s, _ := gqlparser.LoadSchema(&ast.Source{Name: "schema", Input: string(file)})
		ctx.JSON(http.StatusOK, gin.H{"data": intro.New(s)})
	})
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql1", func(ctx *gin.Context) {
		iface := graphql.NewInterface(graphql.InterfaceConfig{Name: "Character", Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: &graphql.NonNull{OfType: graphql.String},
			},
		}})
		input := graphql.NewInputObject(graphql.InputObjectConfig{Name: "LoginInput", Fields: graphql.InputObjectConfigFieldMap{
			"username": &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{OfType: graphql.String},
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{OfType: graphql.String},
			},
		}})
		object := graphql.NewObject(graphql.ObjectConfig{Name: "User", Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: &graphql.NonNull{OfType: graphql.String},
			},
		}, Interfaces: []*graphql.Interface{
			iface,
		}})
		config := graphql.SchemaConfig{Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: &graphql.NonNull{OfType: &graphql.List{OfType: &graphql.NonNull{OfType: object}}},
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: &graphql.NonNull{OfType: input},
					},
				},
			},
		}})}
		schema, _ := graphql.NewSchema(config)
		params := graphql.Params{Schema: schema, RequestString: intro.Query}
		result := graphql.Do(params)
		ctx.JSON(http.StatusOK, result)
	})
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql2", func(ctx *gin.Context) {
		file, _ := os.ReadFile("./assets/intro.json")
		_, _ = ctx.Writer.Write(file)
	})
	_ = r.Run(":8081")
}
