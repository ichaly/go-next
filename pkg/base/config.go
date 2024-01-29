package base

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	App      *App        `mapstructure:"app" jsonschema:"title=App"`
	Oauth    *Oauth      `mapstructure:"oauth" jsonschema:"title=oauth"`
	Cache    *DataSource `mapstructure:"cache" jsonschema:"title=Cache"`
	Database *DataSource `mapstructure:"database" jsonschema:"title=DataSource"`
	Captcha  *Captcha    `mapstructure:"captcha" jsonschema:"title=captcha"`
	Email    *Email      `mapstructure:"email" jsonschema:"title=email"`
	Mobile   *Mobile     `mapstructure:"mobile" jsonschema:"title=mobile"`
	Oss      *Oss        `mapstructure:"oss" jsonschema:"title=oss"`
}

type Oss struct {
	Vendor    string `mapstructure:"vendor"`
	Domain    string `mapstructure:"domain"`
	Bucket    string `mapstructure:"bucket"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

type App struct {
	Name      string `mapstructure:"name" jsonschema:"title=Application Name"`
	Port      string `mapstructure:"port" jsonschema:"title=Application Port"`
	Host      string `mapstructure:"host" jsonschema:"title=Application Host"`
	Debug     bool   `mapstructure:"debug" jsonschema:"title=Debug"`
	Workspace string `mapstructure:"workspace" jsonschema:"title=root"`
}

type Jwt struct {
	Secret string `mapstructure:"secret"`
}

type Oauth struct {
	Jwt      Jwt    `mapstructure:"url"`
	Passkey  string `mapstructure:"passkey"`
	LoginUri string `mapstructure:"login_uri"`
}

type Captcha struct {
	Length  int           `mapstructure:"length" jsonschema:"title=Captcha length"`
	Expired time.Duration `mapstructure:"expired" jsonschema:"title=Captcha expired in minute"`
}

type Email struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	From     string `mapstructure:"from"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Mobile struct {
	SignName        string `mapstructure:"sign_name"`
	TemplateCode    string `mapstructure:"template_code"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
}

type DataSource struct {
	Url      string       `mapstructure:"url"`
	Host     string       `mapstructure:"host"`
	Port     int          `mapstructure:"port"`
	Name     string       `mapstructure:"name"`
	Dialect  string       `mapstructure:"dialect"`
	Username string       `mapstructure:"username"`
	Password string       `mapstructure:"password"`
	Sources  []DataSource `mapstructure:"sources"`
	Replicas []DataSource `mapstructure:"replicas"`
}

func NewConfig(confPath string) (*Config, error) {
	return readInConfig(confPath)
}

func readInConfig(file string) (*Config, error) {
	cp := filepath.Dir(file)
	vi := newViper(cp, filepath.Base(file))

	if err := vi.ReadInConfig(); err != nil {
		return nil, err
	}

	if pcf := vi.GetString("inherits"); pcf != "" {
		cf := vi.ConfigFileUsed()
		vi = newViper(cp, pcf)

		if err := vi.ReadInConfig(); err != nil {
			return nil, err
		}

		if v := vi.GetString("inherits"); v != "" {
			return nil, fmt.Errorf("inherited config (%s) cannot itself inherit (%s)", pcf, v)
		}

		vi.SetConfigFile(cf)

		if err := vi.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "GJ_") || strings.HasPrefix(e, "SJ_") {
			kv := strings.SplitN(e, "=", 2)
			setKeyValue(vi, kv[0], kv[1])
		}
	}

	c := &Config{}

	if err := vi.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to decode config, %v", err)
	}

	return c, nil
}

func newViper(configPath, configFile string) *viper.Viper {
	vi := newViperWithDefaults()
	vi.SetConfigName(strings.TrimSuffix(configFile, filepath.Ext(configFile)))

	if configPath == "" {
		vi.AddConfigPath("./config")
	} else {
		vi.AddConfigPath(configPath)
	}

	return vi
}

func newViperWithDefaults() *viper.Viper {
	vi := viper.New()

	vi.SetDefault("debug", true)

	vi.SetDefault("app.port", "3000")

	vi.SetDefault("oauth.login_uri", "")
	vi.SetDefault("oauth.passkey", "go.next")
	vi.SetDefault("oauth.jwt.secret", "12345")

	vi.SetDefault("captcha.length", 6)
	vi.SetDefault("captcha.expired", "10m")

	vi.SetDefault("cache.dialect", "memory")

	vi.SetDefault("database.dialect", "postgres")
	vi.SetDefault("database.host", "localhost")
	vi.SetDefault("database.port", 5432)
	vi.SetDefault("database.username", "postgres")
	vi.SetDefault("database.password", "")
	vi.SetDefault("database.schema", "public")
	vi.SetDefault("database.pool_size", 10)

	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	vi.AutomaticEnv()

	return vi
}

func setKeyValue(vi *viper.Viper, key string, value interface{}) bool {
	if strings.HasPrefix(key, "GJ_") || strings.HasPrefix(key, "SG_") {
		key = key[3:]
	}
	uc := strings.Count(key, "_")
	k := strings.ToLower(key)

	if vi.Get(k) != nil {
		vi.Set(k, value)
		return true
	}

	for i := 0; i < uc; i++ {
		k = strings.Replace(k, "_", ".", 1)
		if vi.Get(k) != nil {
			vi.Set(k, value)
			return true
		}
	}

	return false
}
