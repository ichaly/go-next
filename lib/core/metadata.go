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

type Record struct {
	DataType          string `gorm:"column:data_type;"`
	IsPrimary         bool   `gorm:"column:is_primary;"`
	IsForeign         bool   `gorm:"column:is_foreign;"`
	IsNullable        bool   `gorm:"column:is_nullable;"`
	TableName         string `gorm:"column:table_name;"`
	ColumnName        string `gorm:"column:column_name;"`
	TableRelation     string `gorm:"column:table_relation;"`
	ColumnRelation    string `gorm:"column:column_relation;"`
	TableDescription  string `gorm:"column:table_description;"`
	ColumnDescription string `gorm:"column:column_description;"`
}

type Table struct {
	Name        string
	Description string
	Columns     map[string]Column
}

type Column struct {
	Type           string `mapstructure:"data_type"`
	Name           string `mapstructure:"column_name"`
	Table          string `mapstructure:"table_name"`
	IsPrimary      bool   `mapstructure:"is_primary"`
	IsForeign      bool   `mapstructure:"is_foreign"`
	IsNullable     bool   `mapstructure:"is_nullable"`
	Description    string `mapstructure:"column_description"`
	TableRelation  string `mapstructure:"table_relation"`
	ColumnRelation string `mapstructure:"column_relation"`
}

func (my Column) SetType(dataType string) Column {
	my.Type = dataType
	return my
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
	var list []*Record
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}
	data := make(map[string]map[string]Column)
	for _, v := range list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(v.ColumnName, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(v.TableName, my.cfg.BlockList...); ok {
			continue
		}

		//解析列
		var c Column
		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:           &c,
			WeaklyTypedInput: true,
			MatchName: func(mapKey, fieldName string) bool {
				mapKey = strcase.ToSnake(mapKey)
				return strings.EqualFold(mapKey, fieldName)
			},
			DecodeHook: func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
				if t != reflect.TypeOf(Column{}) {
					return data, nil
				}
				if val, ok := data.(*Record); !ok {
					return data, nil
				} else {
					if val.IsPrimary {
						val.DataType = "ID"
					} else {
						val.DataType = internal.DataTypes[val.DataType]
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
				Name:        v.TableName,
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

	//构建边信息
	for _, v := range data {
		for f, c := range v {
			pt, pc := my.Named(
				c.Table, c.Name,
				WithTrimSuffix(),
				WithRecursion(c, true),
			)
			ft, fc := my.Named(
				c.TableRelation,
				c.ColumnRelation,
				WithTrimSuffix(),
				PrimaryColumn(pt),
				JoinListSuffix(),
				WithRecursion(c, false),
			)
			//OneToMany
			my.Nodes[pt].Columns[pc] = c.SetType(ft)
			//ManyToOne
			my.Nodes[ft].Columns[fc] = c.SetType(fmt.Sprintf("[%s]", pt))
			//ManyToMany
			rest := maputil.OmitByKeys(v, []string{f})
			for _, s := range rest {
				table, column := my.Named(
					s.TableRelation,
					s.Name,
					WithTrimSuffix(),
					JoinListSuffix(),
				)
				my.Nodes[ft].Columns[column] = s.SetType(fmt.Sprintf("[%s]", table))
			}
		}
	}
	return nil
}
