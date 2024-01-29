package oss

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		//发送验证码
		fx.Annotate(
			NewOss,
			fx.ResultTags(`group:"plugin"`),
		),
		fx.Annotate(
			NewQiNiu,
			fx.ResultTags(`group:"oss"`),
		),
	),
)
