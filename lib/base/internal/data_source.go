package internal

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
