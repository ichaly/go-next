package content

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "schema",
			Target: NewContentQuery,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewContentMutation,
		},
	),
)
