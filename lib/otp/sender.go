package otp

type Sender interface {
	Support(kind string) bool
	Execute(to, code string) error
}
