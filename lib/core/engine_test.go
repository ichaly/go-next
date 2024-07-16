package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/lib/base"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"testing"
)

func TestEngine(t *testing.T) {
	v, err := base.NewViper("../../cfg/dev.yml")
	if err != nil {
		t.Fatal(err)
	}

	args := []interface{}{"postgres", "postgres", "127.0.0.1", 5678, "gcms"}
	d, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s TimeZone=Asia/Shanghai", args...,
	)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		t.Fatal(err)
	}

	m, err := NewMetadata(d, v)
	if err != nil {
		t.Fatal(err)
	}

	e, err := NewEngine(m)
	if err != nil {
		t.Fatal(err)
	}

	s, err := e.Schema()
	if err != nil {
		t.Fatal(err)
	}

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
