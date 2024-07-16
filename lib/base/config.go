package base

import "github.com/spf13/viper"

type Config struct {
	App *App `mapstructure:"app" jsonschema:"title=App"`
	Oss *Oss `mapstructure:"oss" jsonschema:"title=oss"`

	//Oauth   *Oauth        `mapstructure:"oauth" jsonschema:"title=oauth"`
	//Captcha *Captcha      `mapstructure:"captcha" jsonschema:"title=captcha"`
	//Email   *Email        `mapstructure:"email" jsonschema:"title=email"`
	//Mobile  *Mobile       `mapstructure:"mobile" jsonschema:"title=mobile"`
}

type App struct {
	Name      string `mapstructure:"name" jsonschema:"title=Application Name"`
	Port      string `mapstructure:"port" jsonschema:"title=Application Port"`
	Host      string `mapstructure:"host" jsonschema:"title=Application Host"`
	Debug     bool   `mapstructure:"debug" jsonschema:"title=Debug"`
	Workspace string `mapstructure:"workspace" jsonschema:"title=root"`
}

type Oss struct {
	Vendor    string `mapstructure:"vendor"`
	Domain    string `mapstructure:"domain"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

func NewConfig(v *viper.Viper) (*Config, error) {
	c := &Config{}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

//
//type Jwt struct {
//	Secret string `mapstructure:"secret"`
//}
//
//type Oauth struct {
//	Jwt      Jwt    `mapstructure:"url"`
//	Passkey  string `mapstructure:"passkey"`
//	LoginUri string `mapstructure:"login_uri"`
//}
//
//type Captcha struct {
//	Length  int           `mapstructure:"length" jsonschema:"title=Captcha length"`
//	Expired time.Duration `mapstructure:"expired" jsonschema:"title=Captcha expired in minute"`
//}
//
//type Email struct {
//	Port     int    `mapstructure:"port"`
//	Host     string `mapstructure:"host"`
//	From     string `mapstructure:"from"`
//	Username string `mapstructure:"username"`
//	Password string `mapstructure:"password"`
//}
//
//type Mobile struct {
//	SignName        string `mapstructure:"sign_name"`
//	TemplateCode    string `mapstructure:"template_code"`
//	AccessKeyId     string `mapstructure:"access_key_id"`
//	AccessKeySecret string `mapstructure:"access_key_secret"`
//}
