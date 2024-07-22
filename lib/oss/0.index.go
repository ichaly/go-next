package oss

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewOss,
			fx.ParamTags(``, `group:"oss"`),
			fx.ResultTags(`group:"plugin"`),
		),
		fx.Annotate(
			NewQiNiu,
			fx.ResultTags(`group:"oss"`),
		),
		fx.Annotate(
			NewAliYun,
			fx.ResultTags(`group:"oss"`),
		),
		fx.Annotate(
			NewMinio,
			fx.ResultTags(`group:"oss"`),
		),
		fx.Annotate(
			NewLocal,
			fx.ResultTags(`group:"oss"`),
		),
		fx.Annotate(
			NewTencent,
			fx.ResultTags(`group:"oss"`),
		),
	),
)
