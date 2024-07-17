package internal

type PhoneConfig struct {
	CaptchaConfig   `mapstructure:"captcha"`
	SignName        string `mapstructure:"phone.sign_name"`
	TemplateCode    string `mapstructure:"phone.template_code"`
	AccessKeyId     string `mapstructure:"phone.access_key_id"`
	AccessKeySecret string `mapstructure:"phone.access_key_secret"`
}
