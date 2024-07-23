package main

import (
	"context"
	"database/sql"
	"github.com/dosco/graphjin/core/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
	"testing"
)

type _graphJinSuite struct {
	suite.Suite
	db *sql.DB
}

func TestGraphJin(t *testing.T) {
	suite.Run(t, new(_graphJinSuite))
}

func (my *_graphJinSuite) SetupSuite() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5678/gcms?sslmode=disable")
	my.Require().NoError(err)
	my.db = db
}

func (my *_graphJinSuite) TestGraphJin() {
	gj, err := core.NewGraphJin(&core.Config{}, my.db)
	my.Require().NoError(err)

	query := `{sys_user{id}sys_team{id}}`
	ctx := context.WithValue(context.Background(), core.UserIDKey, 1)
	res, err := gj.GraphQL(ctx, query, nil, nil)
	my.Require().NoError(err)
	my.T().Log(res)
}
