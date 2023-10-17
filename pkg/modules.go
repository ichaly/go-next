package pkg

import (
	"github.com/ichaly/go-next/pkg/auth"
	"github.com/ichaly/go-next/pkg/base"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	auth.Modules,
	base.Modules,
)
