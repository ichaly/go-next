package rule

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "schema",
			Target: NewRuleQuery,
		},
		fx.Annotated{
			Group:  "schema",
			Target: NewRuleMutation,
		},
	),
)
