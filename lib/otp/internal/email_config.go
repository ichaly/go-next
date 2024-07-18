package internal

type EmailConfig struct {
	CaptchaConfig `mapstructure:"captcha"`
	Email         `mapstructure:"email"`
}

type Email struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	From     string `mapstructure:"from"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
