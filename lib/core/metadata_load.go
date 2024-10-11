package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

type Node struct {
	Kind        ast.DefinitionKind
	Type        *ast.Type
	Name        string
	Table       string
	Column      string
	Virtual     bool
	Children    []*Node
	Description string
}

type Entry struct {
	DataType          string `gorm:"column:data_type;"`
	Nullable          bool   `gorm:"column:is_nullable;"`
	Iterable          bool   `gorm:"column:is_iterable;"`
	IsPrimary         bool   `gorm:"column:is_primary;"`
	IsForeign         bool   `gorm:"column:is_foreign;"`
	TableName         string `gorm:"column:table_name;"`
	ColumnName        string `gorm:"column:column_name;"`
	TableRelation     string `gorm:"column:table_relation;"`
	ColumnRelation    string `gorm:"column:column_relation;"`
	TableDescription  string `gorm:"column:table_description;"`
	ColumnDescription string `gorm:"column:column_description;"`
}

type Class struct {
	Kind        ast.DefinitionKind
	Name        string
	Table       string
	Fields      map[string]*Field
	Virtual     bool
	Description string
}

type Field struct {
	Name         string
	Type         *ast.Type
	Path         *Entry
	Join         *Entry
	Link         Chain
	Table        string
	Column       string
	Virtual      bool
	Arguments    []*Input
	Description  string
	RelationKind string
}

type Input struct {
	Name        string
	Type        *ast.Type
	Default     *Value
	Description string
}

type Value struct {
}

type Chain string

func (my *Metadata) tableOption() error {
	// 查询表结构
	var list []*Entry
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}

	edge := make(internal.AnyMap[internal.AnyMap[*Entry]])
	//构建节点信息
	for _, r := range list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(r.ColumnName, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(r.TableName, my.cfg.BlockList...); ok {
			continue
		}

		//转化类型
		r.DataType = condition.TernaryOperator(r.IsPrimary, "ID", my.cfg.Mapping[r.DataType])

		//规范命名
		table, column := my.Named(r.TableName, r.ColumnName)

		//构建字段
		field := &Field{
			Type:        ast.NamedType(r.DataType, nil),
			Name:        column,
			Table:       r.TableName,
			Column:      r.ColumnName,
			Description: r.ColumnDescription,
		}

		//索引节点
		maputil.GetOrSet(my.Nodes, table, &Class{
			Name:        table,
			Kind:        ast.Object,
			Table:       r.TableName,
			Description: r.TableDescription,
			Fields:      make(internal.AnyMap[*Field]),
		}).Fields[column] = field

		//索引外键
		if r.IsForeign {
			maputil.GetOrSet(edge, table, make(map[string]*Entry))[column] = r
		}
	}

	//构建关联信息
	for _, v := range edge {
		for k, e := range v {
			currentClass, currentField := my.Named(
				e.TableName, e.ColumnName,
				WithTrimSuffix(),
				NamedRecursion(e, true),
			)
			foreignClass, foreignField := my.Named(
				e.TableRelation,
				e.ColumnRelation,
				WithTrimSuffix(),
				SwapPrimaryKey(currentClass),
				JoinListSuffix(),
				NamedRecursion(e, false),
			)
			//ManyToOne
			my.Nodes[currentClass].Fields[currentField] = &Field{
				Name: currentField,
				Type: ast.NamedType(foreignClass, nil),
				Link: MANY_TO_ONE,
				Path: &Entry{
					TableName:      e.TableRelation,
					ColumnName:     e.ColumnRelation,
					TableRelation:  e.TableName,
					ColumnRelation: e.ColumnName,
				},
				Table:     e.TableName,
				Column:    e.ColumnName,
				Arguments: inputs(foreignClass),
			}
			//OneToMany
			my.Nodes[foreignClass].Fields[foreignField] = &Field{
				Name:      foreignField,
				Type:      ast.ListType(ast.NamedType(currentClass, nil), nil),
				Link:      condition.TernaryOperator(e.TableRelation == e.TableName, RECURSIVE, ONE_TO_MANY),
				Path:      e,
				Arguments: inputs(currentClass),
			}
			//ManyToMany
			rest := maputil.OmitBy(v, func(key string, value *Entry) bool {
				return k == key || value.TableRelation == e.TableName
			})
			for _, r := range rest {
				class, field := my.Named(
					r.TableRelation,
					r.ColumnName,
					WithTrimSuffix(),
					JoinListSuffix(),
				)
				my.Nodes[foreignClass].Fields[field] = &Field{
					Name: field,
					Type: ast.ListType(ast.NamedType(class, nil), nil),
					Link: MANY_TO_MANY,
					Path: &Entry{
						TableName:      r.TableRelation,
						ColumnName:     e.ColumnRelation,
						TableRelation:  e.TableName,
						ColumnRelation: r.ColumnName,
					},
					Join:      e,
					Arguments: inputs(class),
				}
			}
		}
	}

	return nil
}
