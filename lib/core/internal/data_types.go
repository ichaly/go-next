package internal

// DataTypes 内置的数据库到GraphQL的类型映射
var DataTypes = map[string]string{
	"timestamp with time zone": "DateTime",
	"character varying":        "String",
	"text":                     "String",
	"smallint":                 "Int",
	"integer":                  "Int",
	"bigint":                   "Int",
	"smallserial":              "Int",
	"serial":                   "Int",
	"bigserial":                "Int",
	"decimal":                  "Float",
	"numeric":                  "Float",
	"real":                     "Float",
	"double precision":         "Float",
	"money":                    "Float",
	"boolean":                  "Boolean",
}
