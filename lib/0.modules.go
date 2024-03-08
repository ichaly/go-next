package lib

import (
	"github.com/ichaly/go-next/lib/auth"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/oss"
	"github.com/ichaly/go-next/lib/otp"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	auth.Modules,
	base.Modules,
	otp.Modules,
	oss.Modules,
)
