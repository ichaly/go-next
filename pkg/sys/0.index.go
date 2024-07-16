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
				return &Role{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Bind{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &Rule{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &RoleRule{}
			},
		},
		fx.Annotated{
			Group: "entity",
			Target: func() interface{} {
				return &RoleUser{}
			},
		},
		NewUserService,
		NewRoleService,
	),
	fx.Populate(&roleService),
)
