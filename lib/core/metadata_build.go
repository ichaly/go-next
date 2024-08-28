package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/core/internal"
)

type Object struct {
	Name        string
	Description string
	Fields      Directory[Field]
}

type Field struct {
	Type        string `gorm:"column:data_type;"`
	Name        string `gorm:"column:column_name;"`
	Description string `gorm:"column:column_description;"`
}

func (my *Metadata) build() {
	for k, t := range my.tree {
		for f, c := range t.Columns {
			typ := condition.TernaryOperator(c.IsPrimary, "ID", internal.DataTypes[c.Type])

			maputil.GetOrSet(my.Nodes, k, Object{
				Name:        k,
				Description: c.TableDescription,
			}).Fields[f] = Field{
				Name:        f,
				Type:        typ,
				Description: c.Description,
			}

			println(k, f, c.Table, c.Name)
		}
	}

	//构建关联信息
	//for _, v := range my.keys {
	//	for f, c := range v {
	//		currentTable, currentColumn := my.Named(
	//			c.Table, c.Name,
	//			WithTrimSuffix(),
	//			NamedRecursion(c, true),
	//		)
	//		foreignTable, foreignColumn := my.Named(
	//			c.TableRelation,
	//			c.ColumnRelation,
	//			WithTrimSuffix(),
	//			SwapPrimaryKey(currentTable),
	//			JoinListSuffix(),
	//			NamedRecursion(c, false),
	//		)
	//		//OneToMany
	//		my.Nodes[currentTable].Columns[currentColumn] = c.SetType(foreignTable)
	//		//ManyToOne
	//		my.Nodes[foreignTable].Columns[foreignColumn] = c.SetType(fmt.Sprintf("[%s]", currentTable))
	//		if c.Table == c.TableRelation {
	//			continue
	//		}
	//		//ManyToMany
	//		rest := maputil.OmitBy(v, func(key string, value Column) bool {
	//			return f == key || value.TableRelation == c.Table
	//		})
	//		for _, s := range rest {
	//			table, column := my.Named(
	//				s.TableRelation,
	//				s.Name,
	//				WithTrimSuffix(),
	//				JoinListSuffix(),
	//			)
	//			my.Nodes[foreignTable].Columns[column] = s.SetType(fmt.Sprintf("[%s]", table))
	//		}
	//	}
	//}

	//query := Table{Columns: make(map[string]Column)}
	//for k := range my.Nodes {
	//	name := strings.Join([]string{k, "list"}, "_")
	//	if my.cfg.UseCamel {
	//		name = strcase.ToLowerCamel(name)
	//	}
	//	query.Columns[name] = Column{
	//		Name: name, Type: fmt.Sprintf("[%s]", k),
	//	}
	//}
	//my.Nodes["Query"] = query
}
