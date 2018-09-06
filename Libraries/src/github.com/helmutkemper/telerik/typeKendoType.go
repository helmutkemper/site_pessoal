package telerik

type KendoType int

func (el KendoType) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoType) String() string {
	return kendoTypes[el]
}

func (el KendoType) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var kendoTypes = [...]string{
	"",
	"odata",
	"odata-v4",
	"signalr",
}

const (
	KENDO_TYPE_ODATA KendoType = iota + 1
	KENDO_TYPE_ODATA_V4
	KENDO_TYPE_SIGNALR
)
