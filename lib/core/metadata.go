package core

import (
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/jinzhu/inflection"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
	"text/template"
)

//go:embed assets/pgsql.sql
var pgsql string

//go:embed assets/build.tpl
var build string

func init() {
	inflection.AddUncountable("children")
}

type Directory[V any] map[string]V

type Argument struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
}

type Metadata struct {
	db  *gorm.DB
	tpl *template.Template
	cfg *internal.TableConfig

	list []Column
	tree Directory[Table]
	keys Directory[Directory[Column]]

	Nodes Directory[Object]
}

func NewMetadata(v *viper.Viper, d *gorm.DB) (*Metadata, error) {
	cfg := &internal.TableConfig{Mapping: internal.DataTypes}
	if err := v.Sub("schema").Unmarshal(cfg); err != nil {
		return nil, err
	}
	my := &Metadata{
		Nodes: make(Directory[Object]), db: d, cfg: cfg,
		tpl: template.Must(template.New("assets/build.tpl").Funcs(template.FuncMap{
			"toLowerCamel": strcase.ToLowerCamel,
		}).Parse(build)),
	}

	if err := my.load(); err != nil {
		return nil, err
	}
	my.build()

	return my, nil
}

func (my *Metadata) Marshal() (string, error) {
	var w strings.Builder
	if err := my.tpl.Execute(&w, my.Nodes); err != nil {
		return "", err
	}
	return w.String(), nil
}
