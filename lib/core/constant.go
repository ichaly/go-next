package core

// 根结点的名称
const (
	QUERY        = "Query"
	MUTATION     = "Mutation"
	SUBSCRIPTION = "Subscription"
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
