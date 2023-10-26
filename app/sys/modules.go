package sys

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &User{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Team{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Oauth{}
			},
		},
	),
)
