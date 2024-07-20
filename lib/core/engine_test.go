package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/lib/base"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

type EngineSuite struct {
	suite.Suite
	d *gorm.DB
	v *viper.Viper
}

func (my *EngineSuite) SetupSuite() {
	v, err := base.NewViper("../../cfg/dev.yml")
	my.Require().NoError(err)
	d, err := base.NewConnect(v, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	my.Require().NoError(err)
	my.v = v
	my.d = d
}

func (my *EngineSuite) TestEngine() {
	m, err := NewMetadata(my.d, my.v)
	my.Require().NoError(err)

	e, err := NewEngine(m)
	my.Require().NoError(err)

	s, err := e.Schema()
	my.Require().NoError(err)

	r := gin.Default()
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql", func(c *gin.Context) {
		var req struct {
			Query     string                 `form:"query"`
			Operation string                 `form:"operationName" json:"operationName"`
			Variables map[string]interface{} `form:"variables"`
		}
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

func TestEngine(t *testing.T) {
	suite.Run(t, new(EngineSuite))
}
