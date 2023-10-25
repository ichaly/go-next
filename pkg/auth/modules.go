package auth

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		//Casbin鉴权
		NewEnforcer,
		NewSession,
		//Oauth2认证
		NewOauthServer,
		NewOauthTokenStore,
		NewOauthClientStore,
		fx.Annotate(
			NewOauth,
			fx.ResultTags(`group:"plugin"`),
		),
		//Graphql鉴权中间件
		fx.Annotate(
			NewGraphql,
			fx.ResultTags(`group:"middleware"`),
		),
		//手机邮件验证码登录验证
		fx.Annotate(
			NewCaptcha,
			fx.ResultTags(`group:"middleware"`),
		),
		//登录验证中间件
		fx.Annotate(
			NewOauthVerify,
			fx.ResultTags(`group:"middleware"`),
		),
	),
)
