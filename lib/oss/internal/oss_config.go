package internal

type OssConfig struct {
	Vendor    string `mapstructure:"vendor"`
	Domain    string `mapstructure:"domain"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}
