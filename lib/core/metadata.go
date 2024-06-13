package core

import (
	_ "embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"io"
	"regexp"
	"strings"
	"text/tabwriter"
	"text/template"
	"unicode"
)

//go:embed assets/pgsql.sql
var pgsql string

//go:embed assets/schema.tpl
var schema string

var dbTypeRe = regexp.MustCompile(`([a-zA-Z ]+)(\((.+)\))?`)

type Table struct {
	Name        string
	Prefix      string
	Description string
	Columns     []*Column
}

type Column struct {
	Name        string
	Type        string
	Description string
	IsPrimary   bool
	IsForeign   bool
	IsNullable  bool
}

type Metadata struct {
	Nodes map[string]*Table
}

func NewMetadata(db *gorm.DB) (*Metadata, error) {
	var list []*struct {
		Name             string `gorm:"column:column_name;"`
		Type             string `gorm:"column:data_type;"`
		Table            string `gorm:"column:table_name;"`
		IsPrimary        bool   `gorm:"column:is_primary;"`
		IsForeign        bool   `gorm:"column:is_foreign;"`
		IsNullable       bool   `gorm:"column:is_nullable;"`
		Description      string `gorm:"column:column_description;"`
		TableDescription string `gorm:"column:table_description;"`
	}
	if err := db.Raw(pgsql).Scan(&list).Error; err != nil {
		return nil, err
	}
	metadata := &Metadata{
		Nodes: make(map[string]*Table),
	}
	for _, v := range list {
		var c Column
		if err := mapstructure.Decode(v, &c); err != nil {
			return nil, err
		}
		c.Name = strcase.ToCamel(c.Name)
		node, ok := metadata.Nodes[v.Table]
		if !ok {
			name := strcase.ToCamel(v.Table)
			node = &Table{
				Name:        name,
				Columns:     make([]*Column, 0),
				Description: v.TableDescription,
			}
			metadata.Nodes[name] = node
		}
		node.Columns = append(node.Columns, &c)
	}
	return metadata, nil
}

func writeSchema(m *Metadata, out io.Writer) (err error) {
	tmpl, err := template.New("schema").Funcs(template.FuncMap{
		"pascal": toPascalCase, "dbtype": parseDBType,
	}).Parse(schema)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(out, 2, 2, 2, ' ', 0)
	if err = tmpl.Execute(w, m); err != nil {
		return err
	}
	return
}

func toPascalCase(text string) string {
	var sb strings.Builder
	for _, v := range strings.Fields(text) {
		sb.WriteRune(unicode.ToUpper(rune(v[0])))
		sb.WriteString(v[1:])
	}
	return sb.String()
}

func parseDBType(name string) (res [2]string, err error) {
	v := dbTypeRe.FindStringSubmatch(name)
	if len(v) == 4 {
		res = [2]string{v[1], v[3]}
	} else {
		err = fmt.Errorf("invalid db type: %s", name)
	}
	return
}
