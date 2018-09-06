package telerik

type KendoTypeDataJSon int

func (el KendoTypeDataJSon) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoTypeDataJSon) String() string {
	return kendoTypesDataJSon[el]
}

func (el KendoTypeDataJSon) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var kendoTypesDataJSon = [...]string{
	"",
	"json",
	"jsonp",
}

const (
	KENDO_TYPE_DATA_JSON_JSON KendoTypeDataJSon = iota + 1
	KENDO_TYPE_DATA_JSON_JSONP
)
