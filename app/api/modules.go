package api

import (
	"github.com/ichaly/go-next/app/api/content"
	"github.com/ichaly/go-next/app/api/user"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	user.Modules,
	content.Modules,
)
