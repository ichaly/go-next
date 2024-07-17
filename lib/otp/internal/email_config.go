package internal

type EmailConfig struct {
	CaptchaConfig `mapstructure:"captcha"`
	Port          int    `mapstructure:"email.port"`
	Host          string `mapstructure:"email.host"`
	From          string `mapstructure:"email.from"`
	Username      string `mapstructure:"email.username"`
	Password      string `mapstructure:"email.password"`
}
