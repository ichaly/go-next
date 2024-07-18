package internal

type PhoneConfig struct {
	CaptchaConfig `mapstructure:"captcha"`
	Phone         `mapstructure:"phone"`
}

type Phone struct {
	SignName        string `mapstructure:"sign_name"`
	TemplateCode    string `mapstructure:"template_code"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
}
