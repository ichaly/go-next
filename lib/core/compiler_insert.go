package core

import (
	"github.com/duke-git/lancet/v2/convertor"
	queue "github.com/duke-git/lancet/v2/datastructure/queue"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

func (my *compilerContext) renderInsert(id, pid int, f *ast.Field) {
	result := queue.NewLinkedQueue[ast.Value]()
	insert := f.Arguments.ForName(INSERT)
	my.parseValue(insert.Value, result)
	union := make(map[string][]string)

	for !result.IsEmpty() {
		value, _ := result.Dequeue()
		class := strings.TrimSuffix(value.Definition.Name, SUFFIX_INSERT_INPUT)
		table, _ := my.meta.TableName(class, false)
		alias := util.JoinString(table, `_`, convertor.ToString(result.Size()))
		union[table] = append(maputil.GetOrSet(union, table, []string{}), alias)

		my.Quoted(alias)
		my.Write(` AS (INSERT INTO `)
		my.Quoted(table)

		my.Write(` (`)
		for i, v := range value.Children {
			if i != 0 {
				my.Write(`,`)
			}
			field, _ := my.meta.FindField(class, v.Name, false)
			my.Quoted(field.Column)
		}
		my.Write(`) SELECT `)
		for i, v := range value.Children {
			if i != 0 {
				my.Write(`,`)
			}
			if value, err := v.Value.Value(my.variables); err == nil {
				my.Wrap(`'`, value)
				my.Write(`::`)
				//TODO:需要转化为数据库对应的具体类型
				my.Write("text")
			}
		}

		my.Write(` RETURNING `)
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

func (my *compilerContext) parseValue(value *ast.Value, result *queue.LinkedQueue[ast.Value]) {
	result.Enqueue(*value)
	for _, v := range value.Children {
		if v.Value.Definition.Kind == ast.InputObject {
			my.parseValue(v.Value, result)
		}
	}
}
