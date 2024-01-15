package app

import (
	"github.com/ichaly/go-next/app/api"
	"github.com/ichaly/go-next/app/cms"
	"github.com/ichaly/go-next/app/sys"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	api.Modules,
	cms.Modules,
	sys.Modules,
)
