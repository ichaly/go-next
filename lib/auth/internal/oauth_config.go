package internal

type Jwt struct {
	Secret string `mapstructure:"secret"`
}

type OauthConfig struct {
	Jwt      Jwt    `mapstructure:"url"`
	LoginUri string `mapstructure:"login_uri"`
}
