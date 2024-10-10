package core

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
	SUFFIX_EXPRESSION = "Expression"
	SUFFIX_EXPR_LIST  = "ListExpression"
)

// 内置枚举类型
const (
	ENUM_SORT_DIRECTION = "SortDirection"
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
