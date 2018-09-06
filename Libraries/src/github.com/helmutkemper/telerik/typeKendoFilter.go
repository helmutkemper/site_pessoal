package telerik

type KendoFilter int

var kendoFilters = [...]string{
	"",
	"startswith.",
	"endswith",
	"contains",
}

func (el KendoFilter) String() string {
	return kendoFilters[el]
}

const (
	FILTER_STARTS_WITH KendoFilter = iota + 1
	FILTER_ENDS_WITH
	FILTER_CONTAINS
)
