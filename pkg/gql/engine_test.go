package gql

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/lib/sys"
	"net/http"
	"testing"
)

type hello struct {
	SchemaMeta[Query, sys.User] `name:"users"`
}

func (my hello) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}

type gqlRequest struct {
	Query     string                 `form:"query"`
	Operation string                 `form:"operationName" json:"operationName"`
	Variables map[string]interface{} `form:"variables"`
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	err := e.Register(&hello{})
	if err != nil {
		t.Error(err)
	}
	s, err := e.Schema()
	if err != nil {
		t.Error(err)
	}
	r := gin.Default()
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql", func(c *gin.Context) {
		var req gqlRequest
		err := c.ShouldBindBodyWith(&req, binding.JSON)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"errors": gqlerrors.FormatErrors(err)})
			return
		}
		res := graphql.Do(graphql.Params{
			Schema:         s,
			Context:        c.Request.Context(),
			RequestString:  req.Query,
			OperationName:  req.Operation,
			VariableValues: req.Variables,
		})
		c.JSON(http.StatusOK, res)
	})
	_ = r.Run(":8080")
}
