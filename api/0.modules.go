package api

import (
	"github.com/ichaly/go-next/api/content"
	"github.com/ichaly/go-next/api/user"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	user.Modules,
	content.Modules,
)
