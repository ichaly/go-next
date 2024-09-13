package core

import (
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/jinzhu/inflection"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/ast"
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

type Option func(my *Metadata) error

type Metadata struct {
	db  *gorm.DB
	tpl *template.Template
	cfg *internal.TableConfig

	list []Field
	tree internal.AnyMap[Class]
	edge internal.AnyMap[internal.AnyMap[Field]]

	Nodes internal.AnyMap[*ast.Definition]
}

func NewMetadata(v *viper.Viper, d *gorm.DB) (*Metadata, error) {
	cfg := &internal.TableConfig{Mapping: internal.DataTypes}
	if err := v.Sub("schema").Unmarshal(cfg); err != nil {
		return nil, err
	}
	my := &Metadata{
		Nodes: make(internal.AnyMap[*ast.Definition]), db: d, cfg: cfg,
		tpl: template.Must(template.New("assets/build.tpl").Funcs(template.FuncMap{
			"toLowerCamel": strcase.ToLowerCamel,
		}).Parse(build)),
	}

	for _, o := range []Option{
		tableOption,
		buildOption,
	} {
		if err := o(my); err != nil {
			return nil, err
		}
	}

	return my, nil
}

func (my *Metadata) Marshal() (string, error) {
	var w strings.Builder
	if err := my.tpl.Execute(&w, my.Nodes); err != nil {
		return "", err
	}
	return w.String(), nil
}
