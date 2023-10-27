package totp

import (
	"go.uber.org/fx"
)

type DeliverGroup struct {
	fx.In
	All []Deliver `group:"deliver"`
}

type Deliver interface {
	Support(kind string) bool
	Send(code string, to ...string) error
}
