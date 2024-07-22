package otp

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		//发送验证码
		fx.Annotate(
			NewToken,
			fx.ResultTags(`group:"plugin"`),
			fx.ParamTags(``, ``, `group:"sender"`),
		),
		//验证码验证
		fx.Annotate(
			NewTokenVerify,
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewEmail,
			fx.ResultTags(`group:"sender"`),
		),
		fx.Annotate(
			NewPhone,
			fx.ResultTags(`group:"sender"`),
		),
	),
)
