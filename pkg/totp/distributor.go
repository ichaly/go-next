package totp

import (
	"go.uber.org/fx"
)

type DistributorGroup struct {
	fx.In
	All []Distributor `group:"distributor"`
}

type Distributor interface {
	Support(kind string) bool
	Send(code string, to ...string) error
}
