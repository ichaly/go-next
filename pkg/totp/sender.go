package totp

import (
	"go.uber.org/fx"
)

type DeliverGroup struct {
	fx.In
	All []Sender `group:"sender"`
}

type Sender interface {
	Support(kind string) bool
	Send(code string, to string) error
}
