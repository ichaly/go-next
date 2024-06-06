package cms

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Channel{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Media{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Content{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Comment{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Model{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Action{}
			},
		},
	),
)
