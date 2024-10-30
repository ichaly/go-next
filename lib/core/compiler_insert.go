package core

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/util"
	"github.com/samber/lo"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type insertItem struct {
	index  int
	field  *Field
	value  *ast.Value
	parent *Entry
}

func (my *compilerContext) renderInsert(id, pid int, f *ast.Field) {
	insert := f.Arguments.ForName(INSERT)
	result := my.parseValue(insert.Value, nil)
	union := make(map[string][]string)

	//定义CTE进行数据插入
	for index, value := range result {
		class := strings.TrimSuffix(value.value.Definition.Name, SUFFIX_INSERT_INPUT)
		table, _ := my.meta.TableName(class, false)
		alias := util.JoinString(table, `_`, convertor.ToString(index))
		union[table] = append(maputil.GetOrSet(union, table, []string{}), alias)

		children := lo.Filter(lo.Map(value.value.Children, func(item *ast.ChildValue, index int) insertItem {
			field, _ := my.meta.FindField(class, item.Name, false)
			return insertItem{index: index, field: field, value: item.Value}
		}), func(item insertItem, index int) bool {
			return item.field != nil && item.field.Kind == NONE
		})

		my.Quoted(alias)
		my.Space(`AS (INSERT INTO`)
		my.Quoted(table)

		my.Write(` (`)
		for i, v := range children {
			if i != 0 {
				my.Write(`,`)
			}
			my.Quoted(v.field.Column)
		}
		if value.parent != nil {
			my.Write(`,`)
			my.Quoted(value.parent.ColumnName)
		}
		my.Write(`) SELECT `)
		for i, v := range children {
			if i != 0 {
				my.Write(`,`)
			}
			raw, _ := v.value.Value(my.variables)
			my.Wrap(`'`, raw)
			my.Write(`::`)
			my.Write("text") //TODO:需要转化为数据库对应的具体类型
		}
		if value.parent != nil {
			from := util.JoinString(value.parent.TableRelation, `_0`)
			my.Write(`,`)
			my.Quoted(from)
			my.Write(`.`)
			my.Quoted(value.parent.ColumnRelation)
			my.Space(`FROM`)
			my.Quoted(from)
		}

		my.Space(`RETURNING`)
		my.Quoted(table)
		my.Write(`.* ),`)
	}

	//将所有关联的表最后拼接成一个和原表同名的临时表
	keys := maputil.Keys(union)
	for i, k := range keys {
		if i != 0 {
			my.Write(`,`)
		}
		my.Quoted(k)
		my.Space(`AS (`)
		for j, v := range union[k] {
			if j != 0 {
				my.Space(`UNION ALL`)
			}
			my.Space(`SELECT * FROM`)
			my.Quoted(v)
		}
		my.Write(`)`)
	}
}

func (my *compilerContext) parseValue(value *ast.Value, parent *Entry) (result []*insertItem) {
	result = append(result, &insertItem{value: value, parent: parent})
	for _, v := range value.Children {
		if v.Value.Definition.Kind == ast.InputObject {
			var link *Entry
			class := strings.TrimSuffix(value.Definition.Name, SUFFIX_INSERT_INPUT)
			field, _ := my.meta.FindField(class, v.Name, false)
			if field != nil && field.Link != nil && field.Kind != MANY_TO_MANY {
				link = field.Link
			}
			if v.Value.Kind == ast.ListValue {
				for _, c := range v.Value.Children {
					result = append(result, my.parseValue(c.Value, link)...)
				}
			} else {
				result = append(result, my.parseValue(v.Value, link)...)
			}
		}
	}
	return
}
