package pkg

import (
	"github.com/ichaly/go-next/pkg/cms"
	"github.com/ichaly/go-next/pkg/sys"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	cms.Modules,
	sys.Modules,
)
