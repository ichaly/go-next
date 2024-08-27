package core

import (
	_ "embed"
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/jinzhu/inflection"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"reflect"
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

type Table struct {
	Name        string
	Description string
	Columns     map[string]Column
	Args        []Argument
}

type Column struct {
	Type             string `gorm:"column:data_type;"`
	Name             string `gorm:"column:column_name;"`
	Table            string `gorm:"column:table_name;"`
	IsPrimary        bool   `gorm:"column:is_primary;"`
	IsForeign        bool   `gorm:"column:is_foreign;"`
	IsNullable       bool   `gorm:"column:is_nullable;"`
	Description      string `gorm:"column:column_description;"`
	TableRelation    string `gorm:"column:table_relation;"`
	ColumnRelation   string `gorm:"column:column_relation;"`
	TableDescription string `gorm:"column:table_description;"`
}

func (my Column) SetType(dataType string) Column {
	my.Type = dataType
	return my
}

type Argument struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
}

type Metadata struct {
	db    *gorm.DB
	cfg   *internal.TableConfig
	tpl   *template.Template
	Nodes map[string]Table
}

func NewMetadata(v *viper.Viper, d *gorm.DB) (*Metadata, error) {
	cfg := &internal.TableConfig{Mapping: internal.DataTypes}
	if err := v.Sub("schema").Unmarshal(cfg); err != nil {
		return nil, err
	}
	my := &Metadata{
		Nodes: make(map[string]Table), db: d, cfg: cfg,
		tpl: template.Must(template.New("assets/build.tpl").Funcs(template.FuncMap{
			"toLowerCamel": strcase.ToLowerCamel,
		}).Parse(build)),
	}
	if err := my.load(); err != nil {
		return nil, err
	}
	return my, nil
}

func (my *Metadata) MarshalSchema() (string, error) {
	var w strings.Builder
	if err := my.tpl.Execute(&w, my.Nodes); err != nil {
		return "", err
	}
	return w.String(), nil
}

func (my *Metadata) load() error {
	var list []*Column
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}
	data := make(map[string]map[string]Column)
	for _, v := range list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(v.Name, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(v.Table, my.cfg.BlockList...); ok {
			continue
		}

		//解析列
		var c Column
		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result: &c, WeaklyTypedInput: true,
			DecodeHook: func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
				if t != reflect.TypeOf(Column{}) {
					return data, nil
				}
				if val, ok := data.(*Column); !ok {
					return data, nil
				} else {
					if val.IsPrimary {
						val.Type = "ID"
					} else {
						val.Type = internal.DataTypes[val.Type]
					}
					return val, nil
				}
			},
		})
		if err := decoder.Decode(v); err != nil {
			return err
		}

		//解析表
		table, column := my.Named(c.Table, c.Name)

		//索引节点
		if node, ok := my.Nodes[table]; ok {
			node.Columns[column] = c
		} else {
			my.Nodes[table] = Table{
				Name:        v.Table,
				Description: v.TableDescription,
				Columns:     map[string]Column{column: c},
			}
		}

		//索引外键
		if c.IsForeign {
			if _, ok := data[table]; ok {
				data[table][column] = c
			} else {
				data[table] = map[string]Column{column: c}
			}
		}
	}

	//构建关联信息
	for _, v := range data {
		for f, c := range v {
			currentTable, currentColumn := my.Named(
				c.Table, c.Name,
				WithTrimSuffix(),
				WithRecursion(c, true),
			)
			foreignTable, foreignColumn := my.Named(
				c.TableRelation,
				c.ColumnRelation,
				WithTrimSuffix(),
				PrimaryColumn(currentTable),
				JoinListSuffix(),
				WithRecursion(c, false),
			)
			//OneToMany
			my.Nodes[currentTable].Columns[currentColumn] = c.SetType(foreignTable)
			//ManyToOne
			my.Nodes[foreignTable].Columns[foreignColumn] = c.SetType(fmt.Sprintf("[%s]", currentTable))
			if c.Table == c.TableRelation {
				continue
			}
			//ManyToMany
			rest := maputil.OmitBy(v, func(key string, value Column) bool {
				return f == key || value.TableRelation == c.Table
			})
			for _, s := range rest {
				table, column := my.Named(
					s.TableRelation,
					s.Name,
					WithTrimSuffix(),
					JoinListSuffix(),
				)
				my.Nodes[foreignTable].Columns[column] = s.SetType(fmt.Sprintf("[%s]", table))
			}
		}
	}

	query := Table{Columns: make(map[string]Column)}
	for k := range my.Nodes {
		name := strings.Join([]string{k, "list"}, "_")
		if my.cfg.UseCamel {
			name = strcase.ToLowerCamel(name)
		}
		query.Columns[name] = Column{
			Name: name, Type: fmt.Sprintf("[%s]", k),
		}
	}
	my.Nodes["Query"] = query
	return nil
}
