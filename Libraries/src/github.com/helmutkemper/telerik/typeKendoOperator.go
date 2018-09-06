package telerik

type KendoOperator int

var kendoOperators = [...]string{
	"",
	"eq",
	"neq",
	"isnull",
	"isnotnull",
	"lt",
	"lte",
	"gt",
	"gte",
	"startswith",
	"endswith",
	"contains",
	"doesnotcontain",
	"isempty",
	"isnotempty",
}

func (el KendoOperator) String() string {
	return kendoOperators[el]
}

const (
	OPERATOR_EQUAL_TO KendoOperator = iota + 1
	OPERATOR_NOT_EQUAL_TO
	OPERATOR_IS_NULL
	OPERATOR_IS_NOT_NULL
	OPERATOR_LESS_THAN
	OPERATOR_LESS_THAN_OR_EQUAL_TO
	OPERATOR_GREATER_THAN
	OPERATOR_GREATER_THAN_OR_EQUAL_TO
	OPERATOR_STARTS_WITH
	OPERATOR_ENDS_WITH
	OPERATOR_CONTAINS
	OPERATOR_DOES_NOT_CONTAIN
	OPERATOR_IS_EMPTY
	OPERATOR_IS_NOT_EMPTY
)
