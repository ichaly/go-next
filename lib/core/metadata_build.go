package core

import (
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
)

var buildOption = func(my *Metadata) error {
	//构建节点信息
	for t, c := range my.tree {
		for k, f := range c.Fields {
			maputil.GetOrSet(my.Nodes, t, Class{
				Name:        f.Name,
				Description: f.Description,
				Fields:      make(Map[Field]),
			}).Fields[k] = f
		}
	}
	//构建关联信息
	for _, v := range my.keys {
		for f, c := range v {
			currentTable, currentColumn := my.Named(
				c.TableName, c.Name,
				WithTrimSuffix(),
				NamedRecursion(c, true),
			)
			foreignTable, foreignColumn := my.Named(
				c.TableRelation,
				c.ColumnRelation,
				WithTrimSuffix(),
				SwapPrimaryKey(currentTable),
				JoinListSuffix(),
				NamedRecursion(c, false),
			)
			//OneToMany
			my.Nodes[currentTable].Fields[currentColumn] = c.SetType(foreignTable)
			//ManyToOne
			my.Nodes[foreignTable].Fields[foreignColumn] = c.SetType(fmt.Sprintf("[%s]", currentTable))
			if c.TableName == c.TableRelation {
				continue
			}
			//ManyToMany
			rest := maputil.OmitBy(v, func(key string, value Field) bool {
				return f == key || value.TableRelation == c.TableName
			})
			for _, s := range rest {
				table, column := my.Named(
					s.TableRelation,
					s.Name,
					WithTrimSuffix(),
					JoinListSuffix(),
				)
				my.Nodes[foreignTable].Fields[column] = s.SetType(fmt.Sprintf("[%s]", table))
			}
		}
	}
	println("build metadata")
	//query := Class{Fields: make(map[string]Field)}
	//for k := range my.Nodes {
	//	name := strings.Join([]string{k, "list"}, "_")
	//	if my.cfg.UseCamel {
	//		name = strcase.ToLowerCamel(name)
	//	}
	//	query.Fields[name] = Field{
	//		Name: name, Type: fmt.Sprintf("[%s]", k),
	//	}
	//}
	//my.Nodes["Query"] = query
	//sys_edge->user_id->sys_user->id
	//sys_edge->team_id->sys_team->id
	//sys_team->area_id->sys_area->id
	return nil
}
