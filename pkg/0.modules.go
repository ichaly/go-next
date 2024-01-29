package pkg

import (
	"github.com/ichaly/go-next/pkg/auth"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/oss"
	"github.com/ichaly/go-next/pkg/otp"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	auth.Modules,
	base.Modules,
	otp.Modules,
	oss.Modules,
)
