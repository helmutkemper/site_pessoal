package telerik

type KendoTypeData int

func (el KendoTypeData) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoTypeData) String() string {
	return kendoTypesData[el]
}

func (el KendoTypeData) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var kendoTypesData = [...]string{
	"",
	"xml",
	"json",
}

const (
	KENDO_TYPE_DATA_XML KendoTypeData = iota + 1
	KENDO_TYPE_DATA_JSON
)
