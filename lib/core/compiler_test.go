package core

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

type _CompilerSuite struct {
	_MetadataSuite
	m *Metadata
}

func TestCompiler(t *testing.T) {
	suite.Run(t, new(_CompilerSuite))
}

func (my *_CompilerSuite) SetupSuite() {
	my._MetadataSuite.SetupSuite()

	var err error
	my.m, err = NewMetadata(my.v, my.d)
	my.Require().NoError(err)
}

func (my *_CompilerSuite) TestCompiler() {
	c, err := NewCompiler(my.m, my.d)
	my.Require().NoError(err)
	query := "{user{id}team{id}}"
	_, _ = c.Compile(query)
}

func (my *_CompilerSuite) TestIntrospection() {
	c, err := NewCompiler(my.m, my.d)
	my.Require().NoError(err)

	r := gin.Default()
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": c.Introspection()})
	})
	r.POST("/graphql0", func(ctx *gin.Context) {
		file, _ := os.ReadFile("./assets/intro.json")
		_, _ = ctx.Writer.Write(file)
	})

	_ = r.Run()
}
