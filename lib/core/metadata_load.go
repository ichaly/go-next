package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

type Class struct {
	Kind        ast.DefinitionKind
	Name        string
	Table       string
	Fields      map[string]*Field
	Description string
}

type Field struct {
	Type        *ast.Type
	Name        string
	Path        string
	Table       string
	Column      string
	Arguments   []*Input
	Description string
}

type Input struct {
	Name        string
	Type        *ast.Type
	Default     *Value
	Description string
}

type Value struct {
}

func (my *Metadata) tableOption() error {
	// 查询表结构
	var list []*internal.Record
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}

	edge := make(internal.AnyMap[internal.AnyMap[*internal.Record]])
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
			maputil.GetOrSet(edge, table, make(map[string]*internal.Record))[column] = r
		}
	}

	//构建关联信息
	for _, v := range edge {
		for k, r := range v {
			currentTable, currentColumn := my.Named(
				r.TableName, r.ColumnName,
				WithTrimSuffix(),
				NamedRecursion(r, true),
			)
			foreignTable, foreignColumn := my.Named(
				r.TableRelation,
				r.ColumnRelation,
				WithTrimSuffix(),
				SwapPrimaryKey(currentTable),
				JoinListSuffix(),
				NamedRecursion(r, false),
			)
			//OneToMany
			my.Nodes[currentTable].Fields[currentColumn] = &Field{
				Name:      currentColumn,
				Type:      ast.NamedType(foreignTable, nil),
				Arguments: inputs(foreignTable),
			}
			//ManyToOne
			my.Nodes[foreignTable].Fields[foreignColumn] = &Field{
				Name:      foreignColumn,
				Type:      ast.ListType(ast.NamedType(currentTable, nil), nil),
				Arguments: inputs(currentTable),
			}
			//ManyToMany
			rest := maputil.OmitBy(v, func(key string, value *internal.Record) bool {
				return k == key || value.TableRelation == r.TableName
			})
			for _, s := range rest {
				table, column := my.Named(
					s.TableName,
					s.ColumnName,
					WithTrimSuffix(),
					JoinListSuffix(),
				)
				my.Nodes[foreignTable].Fields[column] = &Field{
					Name:      column,
					Type:      ast.ListType(ast.NamedType(table, nil), nil),
					Arguments: inputs(table),
				}
			}
		}
	}

	return nil
}
