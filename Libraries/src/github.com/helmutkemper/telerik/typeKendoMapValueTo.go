package telerik

type KendoMapValueTo int

var kendoMapValueTos = [...]string{
	"",
	"index",
	"dataItem",
}

func (el KendoMapValueTo) String() string {
	return kendoMapValueTos[el]
}

const (
	MAP_VALUE_TO_INDEX KendoMapValueTo = iota + 1
	MAP_VALUE_TO_DATA_ITEM
)
