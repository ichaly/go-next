package core

import (
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
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
	IsPrimary      bool   `mapstructure:"is_primary"`
	IsForeign      bool   `mapstructure:"is_foreign"`
	IsNullable     bool   `mapstructure:"is_nullable"`
	Description    string `mapstructure:"column_description"`
	TableRelation  string `mapstructure:"table_relation"`
	ColumnRelation string `mapstructure:"column_relation"`
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

func (my *Metadata) load() error {
	var list []*Record
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}
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
				if val, ok := data.(Record); !ok {
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
		tableName, columnName := v.TableName, v.ColumnName
		parentTable, parentColumn := v.TableRelation, v.ColumnRelation

		//移除前缀
		if val, ok := util.StartWithAny(tableName, my.cfg.Prefixes...); ok {
			tableName = strings.Replace(tableName, val, "", 1)
		}
		if val, ok := util.StartWithAny(parentTable, my.cfg.Prefixes...); ok {
			parentTable = strings.Replace(parentTable, val, "", 1)
		}

		//表名转大驼峰列转小驼峰
		if my.cfg.UseCamel {
			tableName = strcase.ToCamel(tableName)
			parentTable = strcase.ToCamel(parentTable)
			columnName = strcase.ToLowerCamel(columnName)
			parentColumn = strcase.ToLowerCamel(parentColumn)
		}

		//索引节点
		node, ok := my.Nodes[tableName]
		if !ok {
			node = Table{
				Name:        v.TableName,
				Description: v.TableDescription,
				Columns:     make(map[string]Column),
			}
			my.Nodes[tableName] = node
		}

		parent, ok := my.Nodes[parentTable]
		if !ok {
			parent = Table{
				Name:        v.TableName,
				Description: v.TableDescription,
				Columns:     make(map[string]Column),
			}
			my.Nodes[tableName] = parent
		}
		if c.IsForeign {

		}

		node.Columns[columnName] = c
		parent.Columns[parentColumn] = c
	}
	return nil
}

func (my *Metadata) MarshalSchema() (string, error) {
	var w strings.Builder
	if err := my.tpl.Execute(&w, my.Nodes); err != nil {
		return "", err
	}
	return w.String(), nil
}
