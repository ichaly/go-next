package core

import (
	"database/sql"
	"github.com/dosco/graphjin/core/v3"
	"github.com/gin-gonic/gin"
	"github.com/ichaly/go-next/lib/core/internal/introspection"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
	"net/http"
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
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5678/gcms?sslmode=disable")
	my.Require().NoError(err)
	my.db = db
}

func (my *_GraphJinSuite) TestGraphJin() {
	gj, err := core.NewGraphJin(&core.Config{}, my.db)
	my.Require().NoError(err)

	r := gin.Default()
	r.Match([]string{http.MethodGet, http.MethodPost}, "/graphql", func(ctx *gin.Context) {
		res, _ := gj.GraphQL(ctx, introspection.Query, nil, nil)
		ctx.JSON(http.StatusOK, res)
	})
	_ = r.Run()
}
