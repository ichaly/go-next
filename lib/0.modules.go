package lib

import (
	"github.com/ichaly/go-next/lib/cms"
	"github.com/ichaly/go-next/lib/sys"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	cms.Modules,
	sys.Modules,
)
