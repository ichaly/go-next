package core

import "github.com/samber/lo"

// 根结点的名称
const (
	QUERY        = "Query"
	MUTATION     = "Mutation"
	SUBSCRIPTION = "Subscription"
)

const (
	RECURSIVE    Chain = "Recursive"
	ONE_TO_MANY  Chain = "OneToMany"
	MANY_TO_ONE  Chain = "ManyToOne"
	MANY_TO_MANY Chain = "ManyToMany"
)

// tag名
const (
	TAG_NAME        = "name"
	TAG_DESCRIPTION = "description"
)

// SchemaMeta元数据
const (
	SCHEMA_META = "SchemaMeta"
	PARENT_TYPE = "parentType"
	RESULT_TYPE = "resultType"
)

// GraphQL入参名称后缀
const (
	SUFFIX_SORT_INPUT  = "SortInput"
	SUFFIX_DATA_INPUT  = "DataInput"
	SUFFIX_WHERE_INPUT = "WhereInput"
)

// 路基表达式后缀
const (
	SUFFIX_EXPRESSION      = "Expression"
	SUFFIX_EXPRESSION_LIST = "ListExpression"
)

// 内置枚举类型
const (
	ENUM_IS_INPUT   = "IsInput"
	ENUM_SORT_INPUT = "SortInput"
)

// 内置枚举类型
const (
	SCALAR_ID        = "ID"
	SCALAR_INT       = "Int"
	SCALAR_FLOAT     = "Float"
	SCALAR_STRING    = "String"
	SCALAR_CURSOR    = "Cursor"
	SCALAR_BOOLEAN   = "Boolean"
	SCALAR_DATE_TIME = "DateTime"
)

// 过滤操作符号描述
const (
	descIn                 = "Is in list of values"
	descIs                 = "Is value null (true) or not null (false)"
	descEqual              = "Equals value"
	descNotEqual           = "Does not equal value"
	descGreaterThan        = "Is greater than value"
	descGreaterThanOrEqual = "Is greater than or equal to value"
	descLessThan           = "Is less than value"
	descLessThanOrEqual    = "Is less than or equal to value"
	descLike               = "Value matching pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"
	descILike              = "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"
	descRegex              = "Value matching regular pattern"
	descIRegex             = "Value matching (case-insensitive) regex pattern"
)

// 逻辑关系操作符常量
const (
	NOT = "not"
	AND = "and"
	OR  = "or"
)

const (
	IS      = "is"
	EQ      = "eq"
	IN      = "in"
	GT      = "gt"
	GE      = "ge"
	LT      = "lt"
	LE      = "le"
	NE      = "ne"
	LIKE    = "like"
	I_LIKE  = "iLike"
	REGEX   = "regex"
	I_REGEX = "iRegex"
)

// 顺序不要调整这个会影响内置标量的可用操作符
var operators = []*symbol{
	{IS, "is", descIs},
	{EQ, "=", descEqual},
	{IN, "in", descIn},
	{GT, ">", descGreaterThan},
	{GE, ">=", descGreaterThanOrEqual},
	{LT, "<", descLessThan},
	{LE, "<=", descLessThanOrEqual},
	{NE, "!=", descNotEqual},
	{LIKE, "like", descLike},
	{I_LIKE, "ilike", descILike},
	{REGEX, "~", descRegex},
	{I_REGEX, "~*", descIRegex},
}

// 构建操作符和内置标量的关系
var symbols = map[string][]*symbol{
	SCALAR_ID:      operators[1:7], //[eq,in,gt,ge,lt,le]
	SCALAR_INT:     operators[:8],  //[is,eq,in,gt,ge,lt,le,ne]
	SCALAR_FLOAT:   operators[:8],  //[is,eq,in,gt,ge,lt,le,ne]
	SCALAR_STRING:  operators,      //[is,eq,in,gt,ge,lt,le,ne,like,iLike,regex,iRegex]
	SCALAR_BOOLEAN: operators[1:3], //[is,eq]
}

// 运算符按照名字索引字典
var dictionary = lo.Reduce(operators, func(agg map[string]*symbol, item *symbol, index int) map[string]*symbol {
	agg[item.Name] = item
	return agg
}, map[string]*symbol{})

// 内置标量类型集合
var scalars = []string{SCALAR_ID, SCALAR_INT, SCALAR_FLOAT, SCALAR_STRING, SCALAR_BOOLEAN}
