package internal

import "time"

type CaptchaConfig struct {
	Length  int           `mapstructure:"length" jsonschema:"title=Captcha length"`
	Expired time.Duration `mapstructure:"expired" jsonschema:"title=Captcha expired in minute"`
}
