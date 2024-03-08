package otp

import (
	"go.uber.org/fx"
)

type DeliverGroup struct {
	fx.In
	All []Sender `group:"sender"`
}

type Sender interface {
	Support(kind string) bool
	Execute(to, code string) error
}
