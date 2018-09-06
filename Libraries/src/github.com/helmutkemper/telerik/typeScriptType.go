package telerik

type HtmlScriptType int

func (el HtmlScriptType) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el HtmlScriptType) String() string {
	return htmlScriptTypes[el]
}

func (el HtmlScriptType) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var htmlScriptTypes = [...]string{
	"",
	"text/x-kendo-template",
	"application/javascript",
}

const (
	SCRIPT_TYPE_KENDO_TEMPLATE HtmlScriptType = iota + 1
	SCRIPT_TYPE_JAVASCRIPT
)
