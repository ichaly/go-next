package base

import (
	"github.com/ichaly/go-next/lib/gql"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		gql.NewEngine,
		NewConfig,
		NewServer,
		NewStorage,
		NewValidate,
		fx.Annotate(
			NewConnect,
			fx.ParamTags(``, `group:"gorm"`, `group:"entity"`),
		),
		fx.Annotated{
			Group:  "gorm",
			Target: NewSonyFlake,
		},
		fx.Annotated{
			Group:  "gorm",
			Target: NewCache,
		},
		fx.Annotated{
			Group:  "gorm",
			Target: NewAuditor,
		},
		fx.Annotate(
			NewGraphql,
			fx.As(new(Plugin)),
			fx.ResultTags(`group:"plugin"`),
		),
	),
	fx.Invoke(Bootstrap),
)
